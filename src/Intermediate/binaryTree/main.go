package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	tree := dataStruct.NewBinaryTree(5)
	tree.Root.AddNode(3)
	tree.Root.AddNode(2)
	tree.Root.AddNode(4)
	tree.Root.AddNode(8)
	tree.Root.AddNode(7)
	tree.Root.AddNode(6)
	tree.Root.AddNode(10)
	tree.Root.AddNode(9)

	tree.Print()

	if found, cnt := tree.Search(6); found {
		fmt.Println("found 6 cnt:", cnt)
	} else {
		fmt.Println("not found 6 cnt:", cnt)
	}
}
