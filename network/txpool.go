package network

import (
	"sort"

	"github.com/sharansh123/MyBlockChain/core"
	"github.com/sharansh123/MyBlockChain/types"
)

type TxMapSorter struct {
	transactions []*core.Transaction
}

// Len implements sort.Interface.
func (s *TxMapSorter) Len() int {
	return len(s.transactions)
}

// Less implements sort.Interface.
func (s *TxMapSorter) Less(i int, j int) bool {
	return s.transactions[i].GetFirstSeen() < s.transactions[j].GetFirstSeen()
}

// Swap implements sort.Interface.
func (s *TxMapSorter) Swap(i int, j int) {
	s.transactions[i], s.transactions[j] = s.transactions[j], s.transactions[i]


}

func NewTxMapSorter(txMap map[types.Hash]*core.Transaction) *TxMapSorter {
	txx := make([]*core.Transaction, len(txMap))
	i := 0
	for _, v := range txMap {
		txx[i] = v
		i++
	}
	s := &TxMapSorter{txx}
	sort.Sort(s)
	return s
}

type TxPool struct {
	transactions map[types.Hash]*core.Transaction
}

func NewTxPool() *TxPool {
	return &TxPool{
		transactions: make(map[types.Hash]*core.Transaction),
	}
}

func (p *TxPool) Transactions() []*core.Transaction {
	s := NewTxMapSorter(p.transactions)
	return s.transactions
}

func (p *TxPool) Len() int {
	return len(p.transactions)
}

func (p *TxPool) Flush() {
	p.transactions = make(map[types.Hash]*core.Transaction)
}

func (p *TxPool) Add(tx *core.Transaction) error {
	hash := tx.Hash(core.TxHasher{})
	p.transactions[hash] = tx
	return nil
}

func (p *TxPool) Has(hash types.Hash) bool {
	_, ok := p.transactions[hash]
	return ok
}
