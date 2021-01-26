package main

import "fmt"

func addSlice(a []int, i int) {
	if i < len(a) {
		a = append(a[:i+1], a[i:]...)
		a[i] = 8
	} else {
		a = append(a, 8)
	}
	fmt.Println(a)
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	addSlice(a, 3)

}
