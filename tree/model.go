package tree

import "time"

type Block struct {
	uid uint8
	left *Block
	right *Block
	valid bool
	createdTime time.Time
}


func NewBlock () *Block{
	block := new(Block)
	block.uid = 0
	block.valid = true
	block.createdTime = time.Now()
	return block
}

