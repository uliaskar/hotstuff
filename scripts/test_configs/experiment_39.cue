package config

config: {
	consensus:      "fasthotstuff"
	leaderRotation: "round-robin"
	crypto:         "eddsa"
	communication:  "clique"
	byzantineStrategy: silentproposer: [
		2,
	]
	replicaHosts: [
		"localhost",
	]
	clientHosts: [
		"localhost",
	]
	replicas: 4
	clients:  1
}
