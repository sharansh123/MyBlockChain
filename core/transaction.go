package core

import (
	"fmt"

	"github.com/sharansh123/MyBlockChain/crypto"
)


type Transaction struct{
	Data []byte
	PublicKey crypto.PublicKey
	Signature *crypto.Signature
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) (error){
	sig, err := privKey.Sign(tx.Data)
	if err != nil{
		return err
	}
	tx.PublicKey = privKey.PublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) Verify() (error){
	if tx.Signature == nil {
		return fmt.Errorf("tx has no signature")
	}
	if !tx.Signature.Verify(tx.PublicKey, tx.Data){
		return fmt.Errorf("Invalid Transaction")
	}
	return nil
}