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



// Ehh, oq aconteceu aquia? -> Fri 9 Feb, 21:01
// [user@golang-station the-tree]$ go run main.go
// # the-tree/tree
// tree/model.go:45:21: syntax error: unexpected a, expected (
// tree/model.go:46:4: syntax error: unexpected var at end of statement

func Insert (root *Block, blocks ...*Block){
	for _, block := range blocks{
		if !block.valid || !root.valid{
			return
		}

		if len(block.data) == 0 {
			return
		}

		isValidUID := func() bool{
			var isvalid bool = true
			var warnings uint8
			const criterion = 3
			for _, BYTE := range block.uid{
				if warnings >= criterion{
					break
				}
				for _, BYTE2 := range root.uid{
					if BYTE == BYTE2{
						warnings +=1
						if warnings >= criterion{
							isvalid = false
							break
						}
					}
				}
			}
			return isvalid
		}()




		if !isValidUID{
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
