package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
	val string
}

type StructB struct {
	val int
}

func (a *StructA) String() string {
	return "Val: " + a.val
}

type Printable interface {
	String() string
}

func Println(p Printable) {
	fmt.Println(p.String())
}

func (b *StructB) String() string {
	return "Struct B : " + strconv.Itoa(b.val)
}

func main() {
	a := &StructA{val: "AAA"}
	fmt.Println(a)
	Println(a)
	b := &StructB{val: 1}
	fmt.Println(b)
}
