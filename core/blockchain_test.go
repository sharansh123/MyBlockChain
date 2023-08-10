package core

import (
	"testing"

	"github.com/sharansh123/MyBlockChain/types"
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
	prevHash := getPrevBlockHash(t, bc, uint32(1))
	block := randomBlockWithSignAndPrevHash(t, uint32(1), prevHash)
	assert.Nil(t,bc.AddBlock(block))
	assert.Equal(t,bc.Height(), uint32(1))
}

func TestGetHeader(t *testing.T){
	bc := newBlockChain(t)
	lenBlocks := 10

	for i := 0 ; i < lenBlocks; i++{
		prevHash := getPrevBlockHash(t, bc, uint32(i+1))
		block := randomBlockWithSignAndPrevHash(t, uint32(i+1), prevHash)
		assert.Nil(t, bc.AddBlock(block))
		header, err := bc.GetHeader(uint32(i+1))
		assert.Nil(t,err)
		assert.Equal(t, header, block.Header)
	}

}

func TestAddBlock(t *testing.T){
	bc := newBlockChain(t)
	lenBlocks := 100
	for i := 0; i < lenBlocks; i++{
		prevHash := getPrevBlockHash(t, bc, uint32(i+1))
		block := randomBlockWithSignAndPrevHash(t, uint32(i+1), prevHash)
		assert.Nil(t, bc.AddBlock(block))
		header, _ := bc.GetHeader(uint32(i+1))
		assert.Equal(t, prevHash, header.PrevBlockHash)
	}
}


func getPrevBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash{
	h, err := bc.GetHeader(height-1)
	assert.Nil(t,err)
	return BlockHasher{}.Hash(h)
}