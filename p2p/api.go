package p2p

import (
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
)

const (
	ApiVersion = "0.1"
)

type CCHApi struct {
	svc *ChainChainService
}

func createAPI(s *ChainChainService) interface{} {
	api := &CCHApi{
		svc: s,
	}
	s.api = api
	return api
}

// Peers returns number of peers
func (a *CCHApi) Peers() int {
	return len(a.svc.cch.peers)
}

// GetNodeAddress gets the current node's enode
func (a *CCHApi) GetNodeAddress() string {
	return a.svc.p2p.NodeInfo().Enode
}

// AddPeer attempts to add a peer
func (a *CCHApi) AddPeer(enode string) error {
	n, err := discover.ParseNode(enode)
	if err != nil {
		return err
	}
	a.svc.p2p.AddPeer(n)
	return nil
}

// RemovePeer attempts to remove a peer
func (a *CCHApi) RemovePeer(enode string) error {
	n, err := discover.ParseNode(enode)
	if err != nil {
		return err
	}
	a.svc.p2p.RemovePeer(n)
	return nil
}

// PeerInfo returns information about all current peers
func (a *CCHApi) PeerInfo() []*p2p.PeerInfo {
	return a.svc.p2p.PeersInfo()
}
