package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/ayushn2/distri_vault.git/p2p"
)

func makeServer(listenAddr string, nodes ...string) * FileServer{
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr: listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.DefaultDecoder{},

	}
	tcpTransport:= p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{ 
		EncKey: newEncryptionKey(),
		StorageRoot: listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport: tcpTransport ,
		BootstrapNodes: nodes,
	}
	s := NewFileServer(fileServerOpts)
	tcpTransport.OnPeer = s.OnPeer
	return s 

}

func main(){
	s1 := makeServer(":3000","")
	s2 := makeServer(":4000",":3000")
	s3 := makeServer(":6000",":3000",":4000")

	go func() {log.Fatal(s1.Start())}()
	time.Sleep(500 * time.Millisecond)
	go func() {log.Fatal(s2.Start())}()

	time.Sleep(2 * time.Second)


	go func(){
		log.Fatal(s3.Start())
	}()
	time.Sleep(2 * time.Second)

	for i := 0; i < 20; i++ {
	key := fmt.Sprintf("picture_%d.jpg", i)
	data := bytes.NewReader([]byte("my big data file here"))
	s3.Store(key,data)
	time.Sleep(5 * time.Millisecond)

	if err := s3.store.Delete(key); err != nil{
		log.Fatal(err)
	}
			
	r, err := s3.Get(key)
	if err != nil{
		log.Fatal(err)
	}

	b, err := io.ReadAll(r)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(string(b))
	}

	
}