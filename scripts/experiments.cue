[{
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "round-robin"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            "": []
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "round-robin"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            fork: [4]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "round-robin"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            increaseview: [1]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "round-robin"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            silentproposer: [2]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "round-robin"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            "": []
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "round-robin"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            fork: [4]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "round-robin"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            increaseview: [1]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "round-robin"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            silentproposer: [2]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "fixed"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            "": []
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "fixed"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            fork: [4]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "fixed"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            increaseview: [1]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "fixed"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            silentproposer: [2]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "fixed"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            "": []
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "fixed"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            fork: [4]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "fixed"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            increaseview: [1]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "chainedhotstuff"
        leaderRotation: "fixed"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            silentproposer: [2]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "round-robin"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            "": []
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "round-robin"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            fork: [4]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "round-robin"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            increaseview: [1]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "round-robin"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            silentproposer: [2]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "round-robin"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            "": []
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "round-robin"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            fork: [4]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "round-robin"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            increaseview: [1]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "round-robin"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            silentproposer: [2]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "fixed"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            "": []
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "fixed"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            fork: [4]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "fixed"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            increaseview: [1]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "fixed"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            silentproposer: [2]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "fixed"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            "": []
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "fixed"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            fork: [4]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "fixed"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            increaseview: [1]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "simplehotstuff"
        leaderRotation: "fixed"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            silentproposer: [2]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "round-robin"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            "": []
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "round-robin"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            fork: [4]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "round-robin"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            increaseview: [1]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "round-robin"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            silentproposer: [2]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "round-robin"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            "": []
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "round-robin"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            fork: [4]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "round-robin"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            increaseview: [1]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "round-robin"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            silentproposer: [2]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "fixed"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            "": []
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "fixed"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            fork: [4]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "fixed"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            increaseview: [1]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "fixed"
        crypto:         "ecdsa"
        communication:  "clique"
        byzantineStrategy: {
            silentproposer: [2]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "fixed"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            "": []
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "fixed"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            fork: [4]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "fixed"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            increaseview: [1]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}, {
    config: {
        consensus:      "fasthotstuff"
        leaderRotation: "fixed"
        crypto:         "eddsa"
        communication:  "clique"
        byzantineStrategy: {
            silentproposer: [2]
        }
        replicaHosts: ["localhost"]
        clientHosts: ["localhost"]
        replicas: 4
        clients:  1
    }
}]
