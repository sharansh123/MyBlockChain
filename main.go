package main

import (
	"bytes"
	"fmt"
	"log"
	"time"
	"github.com/sharansh123/MyBlockChain/core"
	"github.com/sharansh123/MyBlockChain/crypto"
	"github.com/sharansh123/MyBlockChain/network"
)

func main(){
	fmt.Println("Hello")
	
	trLocal := network.NewLocalTransport("LOCAL")
	trRemoteA := network.NewLocalTransport("REMOTE_A")
	trRemoteB := network.NewLocalTransport("REMOTE_B")
	trRemoteC := network.NewLocalTransport("REMOTE_C")
	trClient := network.NewLocalTransport("CLIENT")

	trLocal.Connect(trRemoteA)
	trRemoteA.Connect(trRemoteB)
	trRemoteB.Connect(trRemoteC)
	trClient.Connect(trLocal)

	makeRemoteServer([]network.Transport{trRemoteA, trRemoteB, trRemoteC})
	
	go func ()  {
		for{
		//trRemote.SendMessage("LOCAL", []byte("Hello from Remote"))
		if err:= sendTransaction(trClient, trLocal.Addr()); err != nil{
			fmt.Errorf(err.Error())
		}
		time.Sleep(2 * time.Second)
	}
	}()
	privKey := crypto.GeneratePrivateKey()
	localServer := makeServer("LOCAL", trLocal, &privKey)
	localServer.Start()
}

func makeRemoteServer(trs []network.Transport){
	for k,v := range trs {
		id := fmt.Sprintf("REMOTE_%d", k)
		s := makeServer(id, v, nil)
		go s.Start()
	}
}

func makeServer(id string, tr network.Transport, privKey *crypto.PrivateKey) *network.Server {
	opts := network.ServerOpts{
		PrivateKey: privKey,
		ID: id,
		Transports: []network.Transport{tr},
	}
	s, err := network.NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}

	return s
}


func sendTransaction(tr network.Transport, to network.NetAddr) error {
	privKey := crypto.GeneratePrivateKey()
	data := []byte{0x1, 0xa, 0x2, 0xa, 0xb}
	tx := core.NewTransaction(data)
	tx.Sign(privKey)
	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil{
		return err
	}

	msg := network.NewMessage(network.MessageTypeTx,buf.Bytes())
	return tr.SendMessage(to, msg.Bytes())
}