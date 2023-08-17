package core

import (
	"fmt"
	"github.com/sharansh123/MyBlockChain/crypto"
	"github.com/sharansh123/MyBlockChain/types"
)


type Transaction struct{
	Data []byte
	From crypto.PublicKey
	Signature *crypto.Signature
	//cached version of tx hash
	hash types.Hash
}

func NewTransaction(data []byte) *Transaction{
	return &Transaction{
		Data: data,
	}
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) (error){
	sig, err := privKey.Sign(tx.Data)
	if err != nil{
		return err
	}
	tx.From = privKey.PublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) Verify() (error){
	if tx.Signature == nil {
		return fmt.Errorf("tx has no signature")
	}
	if !tx.Signature.Verify(tx.From, tx.Data){
		return fmt.Errorf("invalid Transaction")
	}
	return nil
}

func (tx *Transaction) Hash(h Hasher[*Transaction]) types.Hash{
	if tx.hash.IsZero(){
		tx.hash = h.Hash(tx)
	}
	return tx.hash
}
