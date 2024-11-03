package main

import (
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/ayushn2/distri_vault.git/p2p"
)

type FileServerOpts struct{
	StorageRoot string
	PathTransformFunc PathTransformFunc
	Transport p2p.Transport
	BootstrapNodes []string
}

type FileServer struct{
	FileServerOpts
	
	peerLock sync.Mutex
	peers map[string]p2p.Peer
	store *Store
	quitch chan struct{}
}

func NewFileServer(opts FileServerOpts) *FileServer{
	storeOpts := StoreOpts{
		Root: opts.StorageRoot,
		PathTransformFunc: opts.PathTransformFunc,
	}
	return &FileServer{
		FileServerOpts: opts,
		store: NewStore(storeOpts),
		quitch: make(chan struct{}),
		peers: make(map[string]p2p.Peer),
	}
}

type Payload struct{
	Key string
	Data []byte
}

func ( s *FileServer) StoreData(key string,r io.Reader) error{
	// 1. Store this file to disk
	// 2. Broadcast this file to all known peers in the network
	
	return nil
}

func (s *FileServer) Stop(){
	close(s.quitch)
}

func (s *FileServer) OnPeer(p p2p.Peer) error{
	s.peerLock.Lock()
	defer s.peerLock.Unlock()
	s.peers[p.RemoteAddr().String()] = p

	log.Printf("connected with remote %s", p.RemoteAddr())
	return nil
}

func (s *FileServer) loop(){

	defer func(){
		log.Println("file server stopped user quit action")
		s.Transport.Close()
	}()

	for {
		select{
		case msg := <- s.Transport.Consume():
			fmt.Println(msg)
		case <-s.quitch:
			return
		}
	}
}

func (s *FileServer) bootstrapNetwork() error{
	for _,addr := range s.BootstrapNodes{
		// s.Transport.Dial()
		if len(addr) == 0{
			continue
		}
		go func (addr string){
			fmt.Println("attempting to connect with remote: ",addr)
			if err := s.Transport.Dial(addr); err != nil{
				
				log.Println("dial error: ",err)
			}
		}(addr)
		 
	}
	return nil
}

func (s *FileServer) Start() error{
	if err := s.Transport.ListenAndAccept(); err!= nil{
		return err
	}

	if len(s.BootstrapNodes) != 0{
		s.bootstrapNetwork()
	}
	s.bootstrapNetwork()

	s.loop()

	return  nil
}

















