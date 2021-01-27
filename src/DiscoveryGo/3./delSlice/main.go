package main

import "fmt"

func delSlice(a []int, i int) []int {
	a = append(a[:i], a[i+1:]...)
	fmt.Println(a)
	return a
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	delSlice(a, 3)
}
