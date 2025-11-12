package config

config: {
	consensus:      "simplehotstuff"
	leaderRotation: "fixed"
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
