package core

import (
	"fmt"
	"sync"
	"github.com/go-kit/log"
)


type Blockchain struct{
	logger log.Logger
	headers []*Header
	store Storage
	lock sync.RWMutex
	validator Validator
}

func NewBlockChain(l log.Logger, genesis *Block) (*Blockchain, error){
	bc := &Blockchain{
		headers: []*Header{},
		store: NewMemoryStore(),
		logger: l,
	}

	bc.validator = NewBlockValidator(bc)
	err := bc.addBlockWithoutValidation(genesis)
	return bc, err
}


func (bc *Blockchain) AddBlock(b *Block) error{

	if err := bc.validator.ValidateBlock(b); err != nil{
		return fmt.Errorf("validation failed")
	}
	for _, v := range b.Transactions{
		bc.logger.Log("msg", "executing code in vm" , v.Hash(&TxHasher{}))
		vm := NewVM(v.Data)
		if err := vm.Run(); err != nil{
			return err
		}
	}

	return bc.addBlockWithoutValidation(b)

}

func (bc *Blockchain) GetHeader(height uint32) (*Header, error){
	if height > bc.Height(){
		return nil, fmt.Errorf("Height Bigger than Blockchain's height")
	}

	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return bc.headers[height], nil
}


func (bc *Blockchain) HasBlock(height uint32) bool{
	return height <= bc.Height()
}

func (bc *Blockchain) Height() uint32{
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error{
	bc.lock.Lock()
	defer bc.lock.Unlock()
	bc.headers = append(bc.headers, b.Header)


	bc.logger.Log("msg", "adding new block", "hash", b.Hash(BlockHasher{}), "height", b.Height, "transactions", len(b.Transactions))

	return bc.store.Put(b)

}