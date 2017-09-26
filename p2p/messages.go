package p2p

const (
	handshakeMsg = iota
)

type handshakeMsgData struct {
	Version uint64
	// TODO: Address of node!
	NetworkID uint64
}
