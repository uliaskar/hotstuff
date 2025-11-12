package config

config: {
	consensus:      "fasthotstuff"
	leaderRotation: "fixed"
	crypto:         "ecdsa"
	communication:  "clique"
	byzantineStrategy: "": []
	replicaHosts: [
		"localhost",
	]
	clientHosts: [
		"localhost",
	]
	replicas: 4
	clients:  1
}
