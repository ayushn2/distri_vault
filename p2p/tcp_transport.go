package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represenst the remote node over TCP established  connection
type TCPPeer struct{
	// conn is the underlying connection of the peer
	conn net.Conn

	// if we dial a connection => outbound == true
	// if we accept and retrieve a connection => outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer{
	return &TCPPeer{
		conn : conn,
		outbound: outbound,
	}
}

type TCPTransport struct{
	listenAddress string
	listener net.Listener

	mu sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport{
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error{
	var err error
	
	t.listener, err = net.Listen("tcp",t.listenAddress)
	if err != nil{
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop(){
	for {
		conn, err := t.listener.Accept()
		if err!= nil{
			fmt.Printf("TCP accept error : %s\n",err)
		}

		

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn){
	peer := NewTCPPeer(conn,true)//Outbound peer becoz we are accepting (incoming connection)
	fmt.Printf("new incoming connection %v\n",peer)
}