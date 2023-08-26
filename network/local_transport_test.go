package network

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestConnect(t *testing.T){

	tra := NewLocalTransport("A").(*LocalTransport)
	trb := NewLocalTransport("B").(*LocalTransport)
	tra.Connect(trb)
	trb.Connect(tra)
	assert.Equal(t,tra.peers[trb.addr], trb)
	assert.Equal(t,trb.peers[tra.addr], tra)
}

func TestSend(t *testing.T){
	tra := NewLocalTransport("A").(*LocalTransport)
	trb := NewLocalTransport("B").(*LocalTransport)
	tra.Connect(trb)
	trb.Connect(tra)
	msg := []byte("hello")
	assert.Nil(t,tra.SendMessage(trb.addr,msg))

	rpc := <- trb.Consume()
	b := make([]byte, len(msg))
	rpc.Payload.Read(b)
	assert.Equal(t,"hello",string(b))
}

func TestBroadcast(t *testing.T){
	tra := NewLocalTransport("A").(*LocalTransport)
	trb := NewLocalTransport("B").(*LocalTransport)
	trc := NewLocalTransport("C").(*LocalTransport)
	tra.Connect(trb)
	tra.Connect(trc)
	msg := []byte("Hello")
	assert.Nil(t, tra.Broadcast(msg))
	rpc_b := <- trb.Consume()
	rpc_c := <- trc.Consume()
	out_b := make([]byte, len(msg))
	out_c := make([]byte, len(msg))
	rpc_b.Payload.Read(out_b)
	rpc_c.Payload.Read(out_c)
	assert.Equal(t, msg, out_b)
	assert.Equal(t, msg, out_c)
}