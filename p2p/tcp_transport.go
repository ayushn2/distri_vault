package p2p

import (
	"errors"
	"fmt"
	"log"
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

// Close implements the peer interface
func (p * TCPPeer) Close() error{
	return p.conn.Close()
}

type TCPTransportOpts struct{
	ListenAddr string
	HandshakeFunc HandshakeFunc
	Decoder Decoder 
	OnPeer func(Peer) error
}

type TCPTransport struct{
	TCPTransportOpts
	listener net.Listener
	rpcch chan RPC

	mu sync.RWMutex
	peers map[net.Addr]Peer
}



func NewTCPTransport(opts TCPTransportOpts) *TCPTransport{
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch: make(chan RPC),
	}
}

// Consume implements the Transport interface, whichb will return read-only channel for reading the incoming messages received from another peer
func (t *TCPTransport) Consume() <-chan RPC{
	return t.rpcch
}

// Close implements the Transport interface
func (t *TCPTransport) Close() error{
	return t.listener.Close()
}

// Dial implements the transport interface
func (t *TCPTransport) Dial(addr string) error{
	conn, err := net.Dial("tcp", addr)
	if err!= nil{
		return err
	}

	go t.handleConn(conn, true)

	return nil
}

func (t *TCPTransport) ListenAndAccept() error{
	var err error
	
	t.listener, err = net.Listen("tcp",t.ListenAddr)
	if err != nil{
		return err
	}

	go t.startAcceptLoop()

	log.Printf("TCP transport listening on port  %s\n", t.ListenAddr)

	return nil
}

func (t *TCPTransport) startAcceptLoop(){
	for {
		conn, err := t.listener.Accept()
		if errors.Is(err, net.ErrClosed){
			return
		}
		if err!= nil{
			fmt.Printf("TCP accept error : %s\n",err)
		}

		fmt.Printf("new incoming connection %v\n",conn)

		go t.handleConn(conn, false)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn, outbound bool){

	var err error

	defer func(){
		fmt.Printf("dropping peer connection: %s",err)
		conn.Close()
	}()
	peer := NewTCPPeer(conn,outbound)//Outbound peer becoz we are accepting (incoming connection)

	

	if err = t.HandshakeFunc(peer); err != nil{
		return
	}

	if t.OnPeer != nil{
		if err =  t.OnPeer(peer); err != nil{
			return
		}
	}
	
	// Read loop
	rpc := RPC{}

	for {

		err := t.Decoder.Decode(conn,&rpc)
		if err!=nil{
			return
		}

		rpc.From = conn.RemoteAddr()

		t.rpcch <- rpc

	}

	
}