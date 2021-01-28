package main

import "fmt"

func sortByAscending(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if i != j && a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println(a)
}

func main() {
	a := []int{4, 6, 2, 5, 7, 3, 1, 7}
	sortByAscending(a)
}
