package config

config: {
	consensus:      "simplehotstuff"
	leaderRotation: "fixed"
	crypto:         "ecdsa"
	communication:  "clique"
	byzantineStrategy: increaseview: [
		1,
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
