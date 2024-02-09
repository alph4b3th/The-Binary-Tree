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
	hierarchy uint32
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


func Insert (root *Block, blocks ...*Block){
	for _, block := range blocks{
		if !block.valid || !root.valid{
			return
		}

		if len(block.data) == 0 {
			return
		}

		verify := func VerifyUID(){
			var isvalid bool = true
			var warnings uint8
			const criterion = 3
			for _, BYTE := range block.uid{
				if warning >= criterion{
					break
				}
				for _, BYTE2 := range root.uid{
					if BYTE == BYTE2{
						warning +=1
						if warning >= criterion{
							isvalid = false
							break
						}
				}
			}
			return isvalid
		}

		if !verify() {
			return
		}


		switch {
			case len(block.data) < len(root.data):
				root.left = block
				root.hierarchy += 1

			case len(block.data) > len(root.data):
				root.right = block
				root.hierarchy += 1

			default:
				// ok, ele nao eh menor
				// nem maior...

		}

	}
}

func (block *Block) DefineData(data []byte){
	block.data = data
}

func (block *Block) GetHierarchy() uint32 {
	return block.hierarchy
}
