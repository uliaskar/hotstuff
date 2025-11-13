package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	filePath  = flag.String("file", "hotstuff/local/measurements.json", "Path to Hotstuff measurements.json")
	addr      = flag.String("addr", ":9108", "Listen address (host:port)")
	namespace = flag.String("namespace", "hotstuff", "Metrics namespace")
)

type eventEnvelope struct {
	Type  string          `json:"@type"`
	Event json.RawMessage `json:"Event"`
}

type inner struct {
	ID        *float64 `json:"ID"`
	Client    *bool    `json:"Client"`
	Timestamp *string  `json:"Timestamp"`
}

// collector that reads file on every scrape
type fileCollector struct {
	file      string
	namespace string
	// statics
	eventsInFile *prometheus.Desc
	lastTS       *prometheus.Desc
	lastID       *prometheus.Desc
	p50Gap       *prometheus.Desc
	p95Gap       *prometheus.Desc
	readErrors   *prometheus.Desc
	parseErrors  *prometheus.Desc
}

func newCollector(file, ns string) *fileCollector {
	lblsEv := []string{"event_type", "client"}
	lblsType := []string{"event_type"}

	return &fileCollector{
		file:      file,
		namespace: ns,
		eventsInFile: prometheus.NewDesc(
			prometheus.BuildFQName(ns, "", "events_in_file_total"), "Number of events currenctly present in the measurements file", lblsEv, nil),
		lastTS: prometheus.NewDesc(
			prometheus.BuildFQName(ns, "", "last_event_timestamp_seconds"),
			"Unix timestamp of newest event in file",
			lblsType, nil),
		lastID: prometheus.NewDesc(
			prometheus.BuildFQName(ns, "", "last_event_id"), "ID of the newest event in. the file",
			lblsType, nil),
		p50Gap: prometheus.NewDesc(
			prometheus.BuildFQName(ns, "", "inter_event_seconds_p50"),
			"P50 of inter-event intervals computed from file",
			lblsType, nil),
		p95Gap: prometheus.NewDesc(
			prometheus.BuildFQName(ns, "", "inter_event_seconds_p95"),
			"P95 of inter-event intervals computed from file",
			lblsType, nil),
		readErrors: prometheus.NewDesc(
			prometheus.BuildFQName(ns, "", "read_errors_total"),
			"Scrape-time file read errors",
			nil, nil),
		parseErrors: prometheus.NewDesc(
			prometheus.BuildFQName(ns, "", "parse_errors_total"),
			"Scrape-time JSON parse errors",
			nil, nil),
	}
}

func (c *fileCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.eventsInFile
	ch <- c.lastTS
	ch <- c.lastID
	ch <- c.p50Gap
	ch <- c.p95Gap
	ch <- c.readErrors
	ch <- c.parseErrors
}

func (c *fileCollector) Collect(ch chan<- prometheus.Metric) {
	raw, err := os.ReadFile(c.file)
	if err != nil {
		ch <- prometheus.MustNewConstMetric(c.readErrors, prometheus.GaugeValue, 1)
		return
	}

	events, ok := parseEventsLenient(raw)
	if !ok {
		ch <- prometheus.MustNewConstMetric(c.parseErrors, prometheus.GaugeValue, 1)
		return
	}

	// aggregate
	type key struct {
		t string
		c string // "true"/"false"
	}
	counts := map[key]int{}
	lastTSbyType := map[string]float64{}
	lastIdByType := map[string]float64{}
	tsByType := map[string][]float64{}

	for _, ev := range events {
		evType := normalizeType(ev.Type)

		var in inner
		_ = json.Unmarshal(ev.Event, &in) // best effort

		client := "false"
		if in.Client != nil && *in.Client {
			client = "true"
		}
		counts[key{t: evType, c: client}]++

		// timestamp
		if in.Timestamp != nil {
			if ts, ok := parseRFC3339Sec(*in.Timestamp); ok {
				tsByType[evType] = append(tsByType[evType], ts)
				if ts > lastTSbyType[evType] {
					lastTSbyType[evType] = ts
				}
			}
		}

		// id
		if in.ID != nil {
			if *in.ID > lastIdByType[evType] {
				lastIdByType[evType] = *in.ID
			}
		}
	}

	// emit counts
	for k, v := range counts {
		ch <- prometheus.MustNewConstMetric(c.eventsInFile, prometheus.GaugeValue, float64(v), k.t, k.c)
	}

	// emit last ts / id, and p50/p95 of inter-event gaps
	for t, tsList := range tsByType {
		if len(tsList) > 0 {
			ch <- prometheus.MustNewConstMetric(c.lastTS, prometheus.GaugeValue, lastTSbyType[t], t)
		}
		if id, ok2 := lastIdByType[t]; ok2 {
			ch <- prometheus.MustNewConstMetric(c.lastID, prometheus.GaugeValue, id, t)
		}
		// gaps
		if len(tsList) >= 2 {
			sort.Float64s(tsList)
			gaps := make([]float64, 0, len(tsList)-1)
			for i := 1; i < len(tsList); i++ {
				d := tsList[i] - tsList[i-1]
				if d >= 0 && !math.IsInf(d, 0) && !math.IsNaN(d) {
					gaps = append(gaps, d)
				}
			}
			if len(gaps) > 0 {
				sort.Float64s(gaps)
				ch <- prometheus.MustNewConstMetric(c.p50Gap, prometheus.GaugeValue, quantile(gaps, 0.50), t)
				ch <- prometheus.MustNewConstMetric(c.p95Gap, prometheus.GaugeValue, quantile(gaps, 0.95), t)
			}
		}
	}

	// emit lastID for types with IDs but no timestamps
	for t, id := range lastIdByType {
		if _, seen := tsByType[t]; !seen {
			ch <- prometheus.MustNewConstMetric(c.lastID, prometheus.GaugeValue, id, t)
		}
	}
}

// helpers
var typeSuffixRe = regexp.MustCompile(`([A-Za-z0-9]+)$`)

func normalizeType(full string) string {
	if m := typeSuffixRe.FindStringSubmatch(full); len(m) == 2 {
		return m[1]
	}
	if full == "" {
		return "UnknownEvent"
	}
	return full
}

// accept RFC3339 / RFC3339Nano with or without trailin 'Z'
func parseRFC3339Sec(s string) (float64, bool) {
	if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
		return float64(t.UnixNano()) / 1e9, true
	}
	if strings.HasSuffix(s, "Z") {
		ss := strings.TrimSuffix(s, "Z") + "+00:00"
		if t, err := time.Parse(time.RFC3339Nano, ss); err == nil {
			return float64(t.UnixNano()) / 1e9, true
		}
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return float64(t.Unix()), true
	}
	return 0, false
}

func quantile(sorted []float64, q float64) float64 {
	if len(sorted) == 0 {
		return 0
	}
	if q <= 0 {
		return sorted[0]
	}
	if q >= 1 {
		return sorted[len(sorted)-1]
	}
	pos := q * float64(len(sorted)-1)
	i := int(pos)
	f := pos - float64(i)
	if i+1 < len(sorted) {
		return sorted[i]*(1-f) + sorted[i+1]*f
	}
	return sorted[i]
}

// Accepts either a proper JSON array or newline/comma-separated objects
func parseEventsLenient(b []byte) ([]eventEnvelope, bool) {
	trim := bytes.TrimSpace(b)
	var arr []eventEnvelope

	// already an array
	if len(trim) > 0 && trim[0] == '[' {
		dec := json.NewDecoder(bytes.NewReader(trim))
		dec.UseNumber()
		if err := dec.Decode(&arr); err == nil {
			return arr, true
		}
	}

	//
	s := strings.TrimSpace(string(trim))
	s = strings.TrimSuffix(s, ",")
	wrapped := "[" + s + "]"

	dec := json.NewDecoder(strings.NewReader(wrapped))
	dec.UseNumber()
	if err := dec.Decode(&arr); err == nil {
		return arr, true
	}
	return nil, false
}

func main() {
	flag.Parse()

	collector := newCollector(*filePath, *namespace)
	prometheus.MustRegister(collector)

	http.Handle("/metrics", promhttp.Handler())
	fmt.Printf("Reading %s; serving on %s/metrics\n", *filePath, *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		panic(err)
	}
}
