package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"github.com/sharansh123/MyBlockChain/crypto"
	"github.com/sharansh123/MyBlockChain/types"
)


type Header struct{
	Version uint32
	DataHash types.Hash
	PrevBlockHash types.Hash
	Timestamp int64
	Height 	uint32
}


type Block struct{
	*Header
	Transactions []Transaction
	Validator 	crypto.PublicKey
	Signature	*crypto.Signature
	//Cached version of header hash
	hash types.Hash
}

func NewBlock(h *Header, txx []Transaction) *Block{
	return &Block{
		Header: h,
		Transactions: txx,
	}
}

func (b *Block) Decode(r *io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(b)
}


func (b *Block) Encode(w *io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(b)
}

func (b *Block) Hash(hasher Hasher[*Header]) types.Hash{
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b.Header)
	}
	return b.hash
}

func (b *Block) Sign(privKey crypto.PrivateKey) error{
	sig, err := privKey.Sign(b.Header.Bytes())
	if err != nil {
		return err
	}
	b.Validator = privKey.PublicKey()
	b.Signature = sig
	
	return nil
}

func (b *Block) Verify() error{

	if b.Signature == nil {
		return fmt.Errorf("no sig provided")
	}

	if !b.Signature.Verify(b.Validator, b.Header.Bytes()){
		return fmt.Errorf("invaild Block. Signature Doesn't Match")
	}

	for _, v := range b.Transactions{
		if err := v.Verify(); err != nil{
			return fmt.Errorf("invalid Transaction")
		}
	}

	return nil
}

func (h *Header) Bytes() []byte{
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(h)
	return  buf.Bytes()
}
