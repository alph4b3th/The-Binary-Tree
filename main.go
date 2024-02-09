package main

import "fmt"
import "the-tree/tree"

func init() {}

func main() {
	root := tree.NewBlock()
	block := tree.NewBlock()
	block.data = []byte("tem algo aqui kk")
	tree.Insert(root, block, block)
	fmt.Println(root.GetHierarchy())
}
