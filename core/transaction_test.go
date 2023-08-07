package core

import (
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
	assert.Equal(t, privKey.PublicKey(), tx.PublicKey)
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
	tx.PublicKey = otherPrivKey.PublicKey()
	assert.NotNil(t, tx.Verify())
}