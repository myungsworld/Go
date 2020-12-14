package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	tree := dataStruct.Tree{}
	val := 1
	tree.AddNode(val)
	val++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(val)
			val++
		}
	}
	//재귀호출
	tree.DFS1()
	fmt.Println()
	//스택
	tree.DFS2()
	fmt.Println()
	//큐
	tree.BFS()
}
