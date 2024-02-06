package main

import "fmt"
import "the-tree/tree"

func init(){}

func main(){
	block := tree.NewBlock()
	block2 := tree.NewBlock()
	tree.Insert(block2, block)
	fmt.Println(block)
}
