package main

import "fmt"

type VertexID int
type EdgeID int

func NewVerTextIDGenerator() func() VertexID {
	var next int
	return func() VertexID {
		next++
		return VertexID(next)
	}
}

func NewEdgeIDGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

func NewEdge(eid EdgeID) {
	// ...
}

type runes []rune

func main() {
	// gen := NewVerTextIDGenerator()
	// gen2 := NewEdgeIDGenerator()
	// e := NewEdge(gen())

	var a []rune = runes{65, 66}
	fmt.Println(string(a))
}
