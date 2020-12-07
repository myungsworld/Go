package main

import "fmt"

func main() {
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(4))
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(8))
}

func fibo(x int) int {
	if x == 0 {
		return x
	}
	if x == 1 {
		return x
	}
	// if x >= 2 {
	// 	x = fibo(x-1) + fibo(x-2)
	// 	return x
	// }
	// return 0
	return fibo(x-1) + fibo(x-2)
}
