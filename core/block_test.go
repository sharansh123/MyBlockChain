package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/sharansh123/MyBlockChain/crypto"
	"github.com/sharansh123/MyBlockChain/types"
	"github.com/stretchr/testify/assert"
)


func RandomBlock(height uint32) *Block{
	header := &Header{
		Version: 1,
		Timestamp: time.Now().UnixNano(),
		Height: height,
	}

	return NewBlock(header, []*Transaction{})
}

func randomBlockWithSignAndPrevHash(t *testing.T, height uint32, prevBLockHash types.Hash) *Block{
	privKey := crypto.GeneratePrivateKey()
	bc := RandomBlock(height)
	bc.Header.PrevBlockHash = prevBLockHash
	bc.Transactions = append(bc.Transactions, randomTxWithSignature())
	dataHash, err := CalculateDataHash(bc.Transactions)
	assert.Nil(t, err)
	bc.Header.DataHash = dataHash
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
	dataHash, err := CalculateDataHash(b.Transactions)
	assert.Nil(t,err)
	b.DataHash = dataHash
	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()

	assert.NotNil(t, b.Verify())

}


func TestDecodeEncode(t *testing.T){

	b := randomBlockWithSignAndPrevHash(t, 10, types.Hash{})
	buf := &bytes.Buffer{}
	enc := NewGobBlockEncoder(buf)
	assert.Nil(t, enc.Encode(b))

	dec := NewGobBlockDecoder(buf)

	decodeBlock := &Block{}
	assert.Nil(t, dec.Decode(decodeBlock))

	assert.Equal(t, b, decodeBlock)
}