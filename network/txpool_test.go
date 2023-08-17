package network

import (
	"testing"

	"github.com/sharansh123/MyBlockChain/core"
	"github.com/stretchr/testify/assert"
)

func TestTxPool(t *testing.T){
	p := NewTxPool()
	assert.Equal(t, p.Len(), 0)
}

func TestTxPoolAndTx(t *testing.T){
	p := NewTxPool()
	tx := core.NewTransaction([]byte("foo"))
	p.Add(tx)
	assert.Equal(t, p.Len(), 1)

	p.Flush()
	assert.Equal(t, p.Len(), 0)
}