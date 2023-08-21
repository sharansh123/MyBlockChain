package network

import (
	"fmt"
	"math/rand"
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

func TestSortTransaction(t *testing.T){
	p := NewTxPool()
	txLen := 100
	for i := 0 ; i < txLen; i++{
		s := fmt.Sprint(i)
		tx := core.NewTransaction([]byte("foo"+s))
		tx.SetFirstSeen(int64(rand.Intn((i+1)*1000)))
		assert.Nil(t, p.Add(tx))
	}
	assert.Equal(t, p.Len(), txLen)

	txx := p.Transactions()
	for i := 0 ; i < txLen - 1; i++ {
		assert.True(t, txx[i].GetFirstSeen() < txx[i+1].GetFirstSeen())
	}
}

