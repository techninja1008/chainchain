package p2p

import (
	"log"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
)

// BootNode boots a p2p node for network communication
func BootNode(httpAPIs string, httpPort int, dataDir string) *node.Node {
	stack, err := node.New(&node.Config{
		Name:        "cch",
		Version:     strconv.Itoa(Version),
		DataDir:     dataDir,
		HTTPHost:    "127.0.0.1",
		HTTPPort:    httpPort,
		HTTPModules: strings.Split(httpAPIs, ","),
		P2P: p2p.Config{
			MaxPeers:    20,
			NoDiscovery: true,
			ListenAddr:  "127.0.0.1:0",
		},
	})
	if err != nil {
		log.Fatalf("Failed to create network node: %v", err)
	}

	constructor := func(context *node.ServiceContext) (node.Service, error) {
		return new(ChainChainService), nil
	}

	if err := stack.Register(constructor); err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}

	utils.StartNode(stack)

	return stack
}
