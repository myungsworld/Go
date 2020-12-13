package main

import (
	"dataStruct"
	"fmt"
)

func main() {

	// slice로 만든 stack
	stack := []int{}
	for i := 1; i <= 5; i++ {
		stack = append(stack, i)
	}

	fmt.Println(stack)
	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Println(last)

	}

	queue := []int{}
	for i := 1; i <= 5; i++ {
		queue = append(queue, i)
	}
	fmt.Println(queue)
	for len(queue) > 0 {
		var front int
		front, queue = queue[0], queue[1:]
		fmt.Println(front)

	}
	//////////////////////////////////////////

	//LinkedList로 만든 stack
	stack2 := dataStruct.NewStack()
	for i := 1; i <= 5; i++ {
		stack2.Push(i)
	}

	fmt.Println("NewStack")

	//stack2가 비어질떄까지 돌림
	for !stack2.Empty() {
		val := stack2.Pop()
		fmt.Printf("%d -> ", val)
	}
	fmt.Println()

	queue2 := dataStruct.NewQueue()
	for i := 1; i <= 5; i++ {
		queue2.Push(i)
	}

	fmt.Println("NewQueue")

	for !queue2.Empty() {
		val := queue2.Pop()
		fmt.Printf("%d -> ", val)
	}

}
