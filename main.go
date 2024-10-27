package main

import (
	"fmt"
	"log"

	"github.com/ayushn2/distri_vault.git/p2p"
)

func OnPeer(p2p.Peer) error{
	fmt.Println("doing some logic with the peer outside of the tcp transport")
	return nil
}

func main(){
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.DefaultDecoder{},
		OnPeer: OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func(){
		for {
			msg := <-tr.Consume()
			fmt.Printf("%v\n",msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil{
		log.Fatal(err)
	}

	select {}

	
}