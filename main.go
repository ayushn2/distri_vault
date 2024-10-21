package main

import (
	"log"

	"github.com/ayushn2/distri_vault.git/p2p"
)

func main(){
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil{
		log.Fatal(err)
	}

	select {}

	
}