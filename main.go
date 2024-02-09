package main

import "fmt"
import "the-tree/tree"

func init() {}

func main() {
	root := tree.NewBlock()
	block := tree.NewBlock()
	tree.Insert(root, block)
	fmt.Println(block)
}
