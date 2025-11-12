package config

config: {
	consensus:      "chainedhotstuff"
	leaderRotation: "round-robin"
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
