package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/sharansh123/MyBlockChain/crypto"
	"github.com/stretchr/testify/assert"
)


func RandomBlock(height uint32) *Block{
	header := &Header{
		Version: 1,
		Timestamp: time.Now().UnixNano(),
		Height: height,
	}

	tx := Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(header, []Transaction{tx})
}

func randomBlockWithSignature(t *testing.T, height uint32) *Block{
	privKey := crypto.GeneratePrivateKey()
	bc := RandomBlock(height)
	assert.Nil(t, bc.Sign(privKey))

	return bc
}


func TestHashBlock(t *testing.T){
	b := RandomBlock(10)
	fmt.Println(b.Hash(BlockHasher{}))
}


func TestSignBlock(t *testing.T){
	
	b := RandomBlock(0)

	privKey := crypto.GeneratePrivateKey()

	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()

	assert.NotNil(t, b.Verify())

}