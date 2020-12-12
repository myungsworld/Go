package main

import "fmt"

func main() {
	var a int
	var b int
	var p *int

	p = &a
	a = 3
	b = 2

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)

	p = &b

	fmt.Println(b)
	fmt.Println(p)
	fmt.Println(*p)
}
