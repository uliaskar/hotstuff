package config

config: {
	consensus:      "fasthotstuff"
	leaderRotation: "round-robin"
	crypto:         "ecdsa"
	communication:  "clique"
	byzantineStrategy: fork: [
		4,
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
