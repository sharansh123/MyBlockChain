package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestConnect(t *testing.T){

	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")
	tra.Connect(trb)
	trb.Connect(tra)
	assert.Equal(t,tra.peers[trb.addr], trb)
	assert.Equal(t,trb.peers[tra.addr], tra)
}

func TestSend(t *testing.T){
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")
	tra.Connect(trb)
	trb.Connect(tra)

	assert.Nil(t,tra.SendMessage(trb.addr,[]byte("hello")))

	rpc := <- trb.Consume()
	assert.Equal(t, string(rpc.Payload), "hello")
}