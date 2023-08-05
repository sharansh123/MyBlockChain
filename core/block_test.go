package core

import (
	"bytes"
	"testing"
	"time"
	"github.com/sharansh123/MyBlockChain/types"
	"github.com/stretchr/testify/assert"
)



func TestHeader_Decode_Encode(t *testing.T){
	h := &Header{
		Version: 1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height: 10,
		Nonce: 234,
	}

	buf := &bytes.Buffer{}

	assert.Nil(t, h.EncodeBinary(buf))
	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h, hDecode)
}

func TestBlock_Encode_Decode(t *testing.T){
	b := &Block{
		Header: Header{
			Version: 1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height: 10,
			Nonce: 234,
		},
		Transactions: nil,
	}
	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))
	assert.Equal(t, b, bDecode)
}

func TestBlockHash(t *testing.T){
	b := &Block{
		Header: Header{
			Version: 1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height: 10,
			Nonce: 234,
		},
		Transactions: nil,
	}

	h := b.Hash()

	assert.False(t, h.IsZero())

}