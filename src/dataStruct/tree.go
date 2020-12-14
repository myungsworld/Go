package dataStruct

import "fmt"

type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *TreeNode) AddNode(val int) {
	t.Childs = append(t.Childs, &TreeNode{Val: val})
}

func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}

	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}

func (t *Tree) DFS1() {
	DFS1(t.Root)
}

//재귀호출
func DFS1(node *TreeNode) {
	fmt.Printf("%d -> ", node.Val)

	for i := 0; i < len(node.Childs); i++ {
		DFS1(node.Childs[i])
	}
}

//스택이용
func (t *Tree) DFS2() {
	s := []*TreeNode{}
	s = append(s, t.Root)

	for len(s) > 0 {
		var last *TreeNode
		last, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d -> ", last.Val)
		for i := len(last.Childs) - 1; i >= 0; i-- {
			s = append(s, last.Childs[i])
		}
	}
}

func (t *Tree) BFS() {
	queue := []*TreeNode{}
	queue = append(queue, t.Root)

	for len(queue) > 0 {
		var first *TreeNode
		first, queue = queue[0], queue[1:]

		fmt.Printf("%d->", first.Val)
		for i := 0; i < len(first.Childs); i++ {
			queue = append(queue, first.Childs[i])
		}
	}
}
