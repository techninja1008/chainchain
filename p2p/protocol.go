package p2p

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
)

const (
	Version        = 0
	NetworkID      = 314
	ProtocolLength = uint64(8)
)

// CCH is a wrapper for the protocol
type CCH struct {
	NetworkID uint64
	peers     map[discover.NodeID]*peerContext
	svc       *ChainChainService
}

type peerContext struct {
	p  *p2p.Peer
	rw p2p.MsgReadWriter
}

// Setup initialises defaults
func (c *CCH) Setup() error {
	if c.NetworkID == 0 {
		c.NetworkID = NetworkID
	}

	c.peers = make(map[discover.NodeID]*peerContext)

	return nil
}

// Protocol returns a new instance of the cch protocol
func (c *CCH) Protocol() (p2p.Protocol, error) {
	return p2p.Protocol{
		Name:    "cch",
		Version: Version,
		Length:  ProtocolLength,
		Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
			c.peers[p.ID()] = &peerContext{p, rw}
			return c.loop(c.peers[p.ID()])
		},
	}, nil
}

// Loop initiates the handshake then continually parses messages.
func (c *CCH) loop(p *peerContext) error {
	err := c.handshake(p)
	if err != nil {
		return err
	}

	for !c.svc.stop {
		err := c.handle(p)

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CCH) handshake(p *peerContext) error {
	payload := &handshakeMsgData{
		Version:   uint64(Version),
		NetworkID: uint64(c.NetworkID),
	}

	err := p2p.Send(p.rw, handshakeMsg, payload)
	if err != nil {
		return err
	}

	var msg p2p.Msg
	msg, err = p.rw.ReadMsg()
	if err != nil {
		return err
	}

	if msg.Code != handshakeMsg {
		return fmt.Errorf("first msg has code %x (!= %x)", msg.Code, handshakeMsg)
	}

	var status handshakeMsgData
	if err := msg.Decode(&status); err != nil {
		return fmt.Errorf("<- %v: %v", msg, err)
	}

	if status.NetworkID != c.NetworkID {
		return fmt.Errorf("network id mismatch: %d (!= %d)", status.NetworkID, c.NetworkID)
	}

	// TODO: Handle Downgrade
	if Version != status.Version {
		return fmt.Errorf("protocol version mismatch: %d (!= %d)", status.Version, Version)
	}

	log.Println("Handshake success")

	return nil
}

func (c *CCH) handle(p *peerContext) error {
	packet, err := p.rw.ReadMsg()
	if err != nil {
		return err
	}

	packet.Discard()
	return nil
}
