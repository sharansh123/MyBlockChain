package core

import "fmt"



type Blockchain struct{
	headers []*Header
	store Storage
	validator Validator
}

func NewBlockChain(genesis *Block) (*Blockchain, error){
	bc := &Blockchain{
		headers: []*Header{},
		store: NewMemoryStore(),
	}

	bc.validator = NewBlockValidator(bc)
	err := bc.addBlockWithoutValidation(genesis)
	return bc, err
}


func (bc *Blockchain) AddBlock(b *Block) error{

	if err := bc.validator.ValidateBlock(b); err != nil{
		return fmt.Errorf("validation failed")
	}

	return bc.addBlockWithoutValidation(b)

}


func (bc *Blockchain) HasBlock(height uint32) bool{
	return height <= bc.Height()
}

func (bc *Blockchain) Height() uint32{
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error{
	bc.headers = append(bc.headers, b.Header)
	return bc.store.Put(b)
}