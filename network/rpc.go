package network

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/sharansh123/MyBlockChain/core"
)

type MessageType byte

const (
	MessageTypeTx MessageType = 0x1
	MessageTypeBlock 
)

func NewMessage(t MessageType, data []byte) *Message{
	return &Message{
		Header: t,
		Data: data,
	}
}

func (msg *Message) Bytes() []byte {
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(msg)
	return buf.Bytes()
}

type RPC struct{
	From NetAddr
	Payload io.Reader

}

type DecodedMessage struct {
	From NetAddr
	Data any
}

type RPCDecodeFunc func(RPC) (*DecodedMessage, error)

type RPCProcessor interface{
	ProcessMessage(*DecodedMessage) error
}

type Message struct{
	Header MessageType
	Data []byte 
}


func DefaultPRCDecodeFunc(rpc RPC) ( *DecodedMessage, error) {
	msg := Message{}
	if err := gob.NewDecoder(rpc.Payload).Decode(&msg); err != nil{
		return nil, fmt.Errorf("Failed to decode message from %s: %s", rpc.From, err)
	}

	fmt.Printf("From: " +  string(rpc.From))

	switch msg.Header{
		case MessageTypeTx: 
			tx := new(core.Transaction)
			if err := tx.Decode(core.NewGobTxDecoder(bytes.NewReader(msg.Data))); err != nil{
				return nil,err
			}
			return &DecodedMessage{
				From:rpc.From,
				Data: tx,
			}, nil
		default:
			return nil, fmt.Errorf("Invalid Message Header")
		}
}