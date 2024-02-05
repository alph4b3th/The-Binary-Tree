package tree

import "time"
import "bytes"
import "encoding/gob"
import "crypto/sha256"

type Block struct {
	uid []byte
	left *Block
	right *Block
	valid bool
	createdTime time.Time
}


func NewBlock () *Block{
	block := new(Block)
	block.valid = true
	block.createdTime = time.Now()

	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(block)

	h := sha256.New()
	h.Write(b.Bytes())

	block.uid = h.Sum(b.Bytes())
	return block
}

