package core

import "fmt"


type Validator interface{
	ValidateBlock(*Block) error
}

type BlockValidator struct{
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator{
	return &BlockValidator{
		bc: bc,
	}
}

func (v *BlockValidator) ValidateBlock(b *Block) error{
	
	if v.bc.HasBlock(b.Height){
		return fmt.Errorf("chain already contains block with that height")
	}

	if b.Height != v.bc.Height()+1{
		return fmt.Errorf("Block Height too high")
	}

	if err := b.Verify(); err!=nil{
		return err
	}

	prevHeader, err := v.bc.GetHeader(b.Height-1)

	if err != nil{
		return err
	}

	hash := BlockHasher{}.Hash(prevHeader)

	if hash != b.PrevBlockHash {
		return fmt.Errorf("PrevHash Doesn't match")
	}

	return nil
}