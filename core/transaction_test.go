package core

import (
	"bytes"
	"fmt"
	"testing"
	"github.com/sharansh123/MyBlockChain/crypto"
	"github.com/stretchr/testify/assert"
)


func TestSignTransaction(t *testing.T){
	data := []byte("foo")
	privKey := crypto.GeneratePrivateKey()
	sig, _ := privKey.Sign(data)
	tx := &Transaction{
		Data: data,
	}
	tx.Sign(privKey)
	fmt.Println(*sig)
	assert.Equal(t, privKey.PublicKey(), tx.From)
}

func TestVerifyTx(t *testing.T){
	data := []byte("foo")
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: data,
	}
	tx.Sign(privKey)
	assert.Nil(t,tx.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.From = otherPrivKey.PublicKey()
	assert.NotNil(t, tx.Verify())
}


func randomTxWithSignature() Transaction{
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}
	tx.Sign(privKey)
	return *tx
}


func TestTxEncodeDecode(t *testing.T){
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}
	tx.Sign(privKey)
	buf := &bytes.Buffer{}
	assert.Nil(t,tx.Encode(NewGobTxEncoder(buf)))
	decodeTx := new(Transaction)
	assert.Nil(t, decodeTx.Decode(NewGobTxDecoder(buf)))
	assert.Equal(t, tx, decodeTx)
}