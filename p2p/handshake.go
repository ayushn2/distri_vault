package p2p

import "errors"

// ErrInvalidHandhake is returned if the handshake between the local and remoyte nodes could not be established
var ErrInvalidHandshake = errors.New("invalid handshakes")

// Handshake func is 
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error{
	return nil
}

