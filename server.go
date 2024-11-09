package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

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



func (s *FileServer) broadcast(msg *Message) error{
	peers := []io.Writer{}
	for _, peer := range s.peers {
		peers = append(peers, peer)
	}
	
	mw := io.MultiWriter(peers...)
	return gob.NewEncoder(mw).Encode(msg)
	
}

type Message struct{
	Payload any
}

type MessageStoreFile struct{
	Key string
	Size int
}

func ( s *FileServer) StoreData(key string,r io.Reader) error{
	// 1. Store this file to disk
	// 2. Broadcast this file to all known peers in the network

	var(
		fileBuffer =new(bytes.Buffer)
		tee =  io.TeeReader(r, fileBuffer)
	)

	size, err := s.store.Write(key, tee) 
	if err != nil{
		return err
	}

	msg := Message{
		Payload: MessageStoreFile{
			Key: key,
			Size: int(size),
		},
	}

	msgBuf := new(bytes.Buffer)

	if err := gob.NewEncoder(msgBuf).Encode(msg); err != nil{
		return err
	}

	for _, peer := range s.peers{
		if err:= peer.Send(msgBuf.Bytes()); err != nil{
			return err
		}
	}
	time.Sleep(time.Second * 3)

	
	// Sending big file that needs to be streamed
	for _, peer := range s.peers{
		n, err := io.Copy(peer, fileBuffer)
		if err!=nil{
			return err
		}
		fmt.Println("received and written bytes to disk: ",n)
	}

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

	// fmt.Println(s.Transport.Consume())
		
	
	for {
		
		select{
			
			case rpc := <-s.Transport.Consume():
				
				var msg Message
				if err := gob.NewDecoder(bytes.NewReader(rpc.Payload)).Decode(&msg); err != nil{
					log.Println(err)
					return
				}

				if err := s.handleMessage(rpc.From, &msg); err != nil{
					log.Println(err)
					return
				}


				
			case <-s.quitch:
				return

				
			}

		
			
	}
}

func (s *FileServer) handleMessage(from string, msg *Message) error{
	switch v := msg.Payload.(type){
	case MessageStoreFile:
		return s.handleMessageStoreFile(from, v)
	}
	return nil
}

func (s *FileServer) handleMessageStoreFile(from string, msg  MessageStoreFile) error{
	peer, ok := s.peers[from]
	if !ok{
		return fmt.Errorf("peer (%s) could not be found in the peer list",from)
	}

	n,err := s.store.Write(msg.Key, io.LimitReader(peer,int64(msg.Size)))
	if err != nil{
		return err
	}

	fmt.Printf("written %d bytes to disk\n",n)

	peer.(*p2p.TCPPeer).Wg.Done()

	return nil
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
	
	s.loop()
	

	return  nil
}

func init(){
	gob.Register(MessageStoreFile{})
}















