package main

import "fmt"

//재귀함수
func fac(n int) int {
	if n <= 0 {
		return 1
	}
	return n * fac(n-1)
}

//for문
func facItr(n int) int {
	result := 1
	for n > 0 {
		result *= n
		n--
	}
	return result
}

func facItr2(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(fac(5))
	fmt.Println(facItr(5))
}
