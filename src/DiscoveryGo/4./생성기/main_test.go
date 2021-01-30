package main

import "fmt"

func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

func NewVerTextIDGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

func NewEdgeIDGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

func NewEdge(eid int) {

}

func ExampleNewIntGenerator() {
	gen := NewIntGenerator()
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	// Output:
	// .
}
