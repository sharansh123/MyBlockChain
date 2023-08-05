package main

import (
	"fmt"
	"time"

	"github.com/sharansh123/MyBlockChain/network"
)

func main(){
	fmt.Println("Hello")
	
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")
	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func ()  {
		trRemote.SendMessage("LOCAL", []byte("Hello from Remote"))
		time.Sleep(1 * time.Second)
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}