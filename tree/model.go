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
	data []byte
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


func Insert (block, root *Block){
	if !block.valid || !root.valid{
		return
	}

	if len(block.data) == 0 {
		return
	}

	switch {
		case len(block.data) < len(root.data):
			root.left = block

		case len(block.data) > len(root.data):
			root.right = block

		default:
			// ok, ele nao eh menor
			// nem maior...
	}
}
