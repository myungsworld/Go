package main

import "fmt"

func addOne(nums []int) {
	fmt.Println("func nums[0]:", nums[0])
	for i := range nums {
		nums[i]++
	}
	fmt.Println("func n[0] After addone:", nums[0])
}

func addString(s []string) {
	for i := range s {
		s[i] += "k"
	}
	fmt.Println(s)
}

func main() {

	n := []int{1, 2, 3, 4}
	fmt.Println("Main n[0]:", n[0])
	addOne(n)
	fmt.Println("Main n[0] After addOne:", n[0])

	s := []string{"a", "b", "c", "d", "e"}
	addString(s)
}
