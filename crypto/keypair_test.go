package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sign_Verify_Valid(t *testing.T){
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	//addr := pubKey.Address()

	msg := []byte("Hello")

	sign, err := privKey.Sign(msg)

	assert.True(t, sign.Verify(pubKey, []byte("Hello")))
	assert.Nil(t,err)

}

func Test_Sign_Verify_InValid(t *testing.T){
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	//addr := pubKey.Address()

	msg := []byte("Hello")

	sign, err := privKey.Sign(msg)

	assert.False(t, sign.Verify(pubKey, []byte("hello")))
	assert.Nil(t,err)

}