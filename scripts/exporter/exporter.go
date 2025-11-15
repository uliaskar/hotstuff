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
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	filePath  = flag.String("file", "data/local/measurements.json", "Path to Hotstuff measurements.json")
	addr      = flag.String("addr", ":9108", "Listen address (host:port)")
	namespace = flag.String("namespace", "hotstuff", "Metrics namespace")
)

type eventEnvelope struct {
	Type  string          `json:"@type"`
	Event json.RawMessage `json:"Event"`

	// throughput measurements
	Commits  *string `json:"Commits,omitempty"`
	Commands *string `json:"Commands,omitempty"`
	Duration *string `json:"Duration,omitempty"`
	// latency measurements
	Latency *float64 `json:"Latency,omitempty"`
	//Variance *float64 `json:"Variance,omitempty"`
	Variance *FloatOrNaN `json:"Variance,omitempty"`
	Count    *string     `json:"Count,omitempty"`
}

type inner struct {
	ID        *float64 `json:"ID"`
	Client    *bool    `json:"Client"`
	Timestamp *string  `json:"Timestamp"`
}

type config struct {
	Crypto            string   `json:"Crypto"`
	Consensus         string   `json:"Consensus"`
	LeaderRotation    string   `json:"LeaderRotation"`
	BatchSize         float64  `json:"BatchSize"`
	ConnectTimeout    string   `json:"ConnectTimeout"`
	InitialTimeout    string   `json:"InitialTimeout"`
	MaxTimeout        string   `json:"MaxTimeout"`
	TimeoutSamples    float64  `json:"TimeoutSamples"`
	TimeoutMultiplier float64  `json:"TimeoutMultiplier"`
	ByzantineStrategy string   `json:"ByzantineStrategy"`
	SharedSeed        string   `json:"SharedSeed"`
	Modules           []string `json:"Modules"`
	Locations         []string `json:"Locations"`
	TreePositions     []int    `json:"TreePositions"`
	BranchFactor      float64  `json:"BranchFactor"`
	TreeDelta         string   `json:"TreeDelta"`
}

type throughputMeasurement struct {
	ID        *float64 `json:"ID,omitempty"`
	Client    *bool    `json:"Client,omitempty"`
	Timestamp *string  `json:"Timestamp,omitempty"`

	Throughput float64 `json:"Throughput,omitempty"` // e.g. commands/sec
	Duration   float64 `json:"Duration,omitempty"`   // e.g. seconds
	Count      float64 `json:"Count,omitempty"`      // e.g. number of commands
}

//-------Ulia'skar's Addition Below-------//
type FloatOrNaN struct {
	Valid bool
	Value float64
	Raw   string
}

func (f *FloatOrNaN) UnmarshalJSON(b []byte) error {
	// null → leave as zero-value
	if string(b) == "null" {
		return nil
	}

	// Try as number first
	var num float64
	if err := json.Unmarshal(b, &num); err == nil {
		f.Valid = true
		f.Value = num
		f.Raw = string(b)
		return nil
	}

	// Try as string
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		f.Raw = s

		// Treat "NaN" (or empty) as "no usable value", but not an error
		if s == "NaN" || s == "" {
			return nil
		}

		// If it's a numeric string, parse it
		if v, err2 := strconv.ParseFloat(s, 64); err2 == nil {
			f.Valid = true
			f.Value = v
		}
		return nil
	}

	// If it’s something weird, just ignore without failing the whole decode
	return nil
}

//------------------------------//

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

	configInfo         *prometheus.Desc
	configBatchSize    *prometheus.Desc
	configBranchFactor *prometheus.Desc

	throughput *prometheus.Desc
	latency    *prometheus.Desc
}

func newCollector(file, ns string) *fileCollector {
	lblsEv := []string{"event_type", "client"}
	lblsType := []string{"event_type"}

	return &fileCollector{
		file:      file,
		namespace: ns,
		eventsInFile: prometheus.NewDesc(
			prometheus.BuildFQName(ns, "", "events_in_file_total"), "Number of events currently present in the measurements file", lblsEv, nil),
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
		configInfo: prometheus.NewDesc(
			prometheus.BuildFQName(ns, "", "config_info"),
			"Static configuration of this HotStuff run (labels only; value is always 1)",
			[]string{"crypto", "consensus", "leader_rotation", "byzantine_strategy"},
			nil,
		),
		configBatchSize: prometheus.NewDesc(
			prometheus.BuildFQName(ns, "", "config_batch_size"),
			"Batch size configured for this run",
			nil, nil,
		),
		configBranchFactor: prometheus.NewDesc(
			prometheus.BuildFQName(ns, "", "config_branch_factor"),
			"Branch factor of the tree configured for this run",
			nil, nil,
		),
		throughput: prometheus.NewDesc(
			prometheus.BuildFQName(ns, "", "throughput_ops_per_second"),
			"Throughput reported by ThroughputMeasurement events (ops/sec)",
			[]string{"event_type", "client"},
			nil,
		),
		latency: prometheus.NewDesc(
			prometheus.BuildFQName(ns, "", "latency_seconds"),
			"Latency reported by LatencyMeasurement events (seconds)",
			[]string{"event_type", "client"},
			nil,
		),
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
	ch <- c.configInfo
	ch <- c.configBatchSize
	ch <- c.configBranchFactor
	ch <- c.throughput
	ch <- c.latency
}

func (c *fileCollector) Collect(ch chan<- prometheus.Metric) {
	raw, err := os.ReadFile(c.file)
	if err != nil {
		fmt.Printf("read error opening %s: %v\n", c.file, err) // <-- ADD THIS LINE
		ch <- prometheus.MustNewConstMetric(c.readErrors, prometheus.GaugeValue, 1)
		return
	}

	events, ok := parseEventsLenient(raw)
	if !ok {
		ch <- prometheus.MustNewConstMetric(c.parseErrors, prometheus.GaugeValue, 1)
		return
	}

	var cfg *config

	for _, ev := range events {
		var candidate config
		if err := json.Unmarshal(ev.Event, &candidate); err == nil && candidate.Crypto != "" && candidate.Consensus != "" {
			cfg = &candidate
			break
		}
	}

	if cfg != nil {
		// label-only info metric
		ch <- prometheus.MustNewConstMetric(
			c.configInfo,
			prometheus.GaugeValue,
			1,
			cfg.Crypto,
			cfg.Consensus,
			cfg.LeaderRotation,
			cfg.ByzantineStrategy,
		)

		// numeric configs
		ch <- prometheus.MustNewConstMetric(c.configBatchSize, prometheus.GaugeValue, cfg.BatchSize)
		ch <- prometheus.MustNewConstMetric(c.configBranchFactor, prometheus.GaugeValue, cfg.BranchFactor)
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
	throughputByKey := map[key][]float64{}
	latencyByKey := map[key][]float64{}

	for _, ev := range events {
		evType := normalizeType(ev.Type)

		var in inner
		_ = json.Unmarshal(ev.Event, &in) // best effort

		client := "false"
		if in.Client != nil && *in.Client {
			client = "true"
		}
		counts[key{t: evType, c: client}]++

		if evType == "ThroughputMeasurement" {
			if ev.Commands != nil && ev.Duration != nil {
				cmds, err1 := strconv.ParseFloat(*ev.Commands, 64)
				dur, err2 := time.ParseDuration(*ev.Duration)
				if err1 == nil && err2 == nil && dur > 0 {
					throughput := cmds / dur.Seconds()
					if !math.IsNaN(throughput) && !math.IsInf(throughput, 0) {
						// store, don't emit yet
						k := key{t: evType, c: client}
						throughputByKey[k] = append(throughputByKey[k], throughput)
					}
				}
			}
		}
		if evType == "LatencyMeasurement" {
			if ev.Latency != nil {
				k := key{t: evType, c: client}
				latencyByKey[k] = append(latencyByKey[k], *ev.Latency)
			}
		}

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
	// emit throughput (one sample per label set)
	for k, vals := range throughputByKey {
		if len(vals) == 0 {
			continue
		}

		// choose how to aggregate: avg, max, last, etc.
		// Here we'll use average throughput over all ThroughputMeasurement events.
		var sum float64
		for _, v := range vals {
			sum += v
		}
		avg := sum / float64(len(vals))

		ch <- prometheus.MustNewConstMetric(
			c.throughput,
			prometheus.GaugeValue,
			avg,
			k.t,
			k.c,
		)
	}
	// emit latency (one sample per label set)
	for k, vals := range latencyByKey {
		if len(vals) == 0 {
			continue
		}

		var sum float64
		for _, v := range vals {
			sum += v
		}
		avg := sum / float64(len(vals))

		ch <- prometheus.MustNewConstMetric(
			c.latency,
			prometheus.GaugeValue,
			avg,
			k.t,
			k.c,
		)
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
