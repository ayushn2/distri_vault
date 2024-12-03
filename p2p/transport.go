package p2p

import "net"

// Peer is an interface that represents the remote node
type Peer interface{
	net.Conn
	Send(data []byte) error
	CloseStream()
}

// Transport is anything that handles the communictaion between nodes in the network. This can be of the form (TCP, UDP, websockets, ...)
type Transport interface{
	Addr() string
	Dial(string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}