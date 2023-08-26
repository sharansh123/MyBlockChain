package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/sharansh123/MyBlockChain/core"
	"github.com/sharansh123/MyBlockChain/crypto"
	"github.com/sharansh123/MyBlockChain/network"
)

func main(){
	fmt.Println("Hello")
	
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")
	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	
	go func ()  {
		for{
		//trRemote.SendMessage("LOCAL", []byte("Hello from Remote"))
		if err:= sendTransaction(trRemote, trLocal.Addr()); err != nil{
			fmt.Errorf(err.Error())
		}
		time.Sleep(1 * time.Second)
	}
	}()
	privKey := crypto.GeneratePrivateKey()
	opts := network.ServerOpts{
		PrivateKey: &privKey,
		ID: "LOCAL",
		Transports: []network.Transport{trLocal},
	}

	s, _ := network.NewServer(opts)
	s.Start()
}


func sendTransaction(tr network.Transport, to network.NetAddr) error {
	privKey := crypto.GeneratePrivateKey()
	data := []byte(strconv.FormatInt(int64(rand.Intn(10000000)), 10))
	tx := core.NewTransaction(data)
	tx.Sign(privKey)
	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil{
		return err
	}

	msg := network.NewMessage(network.MessageTypeTx,buf.Bytes())
	return tr.SendMessage(to, msg.Bytes())
}