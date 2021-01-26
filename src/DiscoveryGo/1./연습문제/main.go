package main

import "fmt"

func singing(n int) {
	for i := 1; i < n+1; i++ {
		fmt.Printf("타잔이 %d원짜리 팬티를 입고 %d원짜리 칼을차고 노래를 한다\n", i*10, i*10+10)
	}
}

func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	temp := 0
	p, q := 0, 1

	for i := 2; i < n+1; i++ {
		temp = p + q
		p, q = q, temp
	}

	return temp

}

func Hanoi(n int) {
	fmt.Printf("Number of disks : %d\n", n)
	
}

func main() {
	//fmt.Println(fibo(3))
	//singing(4)
}
