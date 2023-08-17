package core

import (
	"crypto/sha256"
	"github.com/sharansh123/MyBlockChain/types"
)


 type Hasher[T any] interface{
	Hash(T) types.Hash

 }


 type BlockHasher struct{}

 func (BlockHasher) Hash(b *Header) types.Hash{
	h := sha256.Sum256(b.Bytes())
	return h
 }

 type TxHasher struct{}

 func (TxHasher) Hash(tx *Transaction) types.Hash{
	h := types.Hash(sha256.Sum256(tx.Data))
	return h
 }