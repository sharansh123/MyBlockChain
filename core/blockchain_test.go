package core

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func newBlockChain(t *testing.T) *Blockchain{
	bc, err := NewBlockChain(RandomBlock(0))
	assert.Nil(t, err)
	assert.NotNil(t,bc.validator)
	return bc
}


func TestAddBlockchain(t *testing.T){
	
	bc := newBlockChain(t)
	b := randomBlockWithSignature(t,1)

	assert.Nil(t,bc.AddBlock(b))
	assert.Equal(t,bc.Height(), uint32(1))
}