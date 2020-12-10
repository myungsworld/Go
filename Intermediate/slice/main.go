package main

import "fmt"

func main() {
	a := make([]int, 2, 4)

	b := append(a, 3)
	b = append(a, 3)
	b = append(a, 3)
	fmt.Printf(" %p %p ", a, b)

}
