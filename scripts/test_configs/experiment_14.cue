package config

config: {
	consensus:      "chainedhotstuff"
	leaderRotation: "fixed"
	crypto:         "eddsa"
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
