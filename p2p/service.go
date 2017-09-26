package p2p

import (
	"log"

	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
)

// ChainChainService is the main p2p service for chainchain that communicates with other peers
type ChainChainService struct {
	cch  *CCH
	api  *CCHApi
	p2p  *p2p.Server
}

// Protocols returns protocols used by the service
func (s *ChainChainService) Protocols() []p2p.Protocol {
	s.cch = &CCH{
		svc: s,
	}
	s.cch.Setup()

	protocol, err := s.cch.Protocol()
	if err != nil {
		panic(err)
	}

	return []p2p.Protocol{protocol}
}

// APIs Returns RPC APIs
func (s *ChainChainService) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "cch",
			Version:   ApiVersion,
			Public:    false,
			Service:   createAPI(s),
		},
	}
}

// Start starts the service
func (s *ChainChainService) Start(p *p2p.Server) error {
	log.Println("Service starting...")
	log.Println("Storing p2p server...")
	s.p2p = p

	log.Printf("Listening on %s for p2p\n", p.ListenAddr)
	return nil
}

// Stop stops the service
func (s *ChainChainService) Stop() error {
	log.Println("Service stopping...")
	return nil
}
