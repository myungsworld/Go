package main

import "fmt"

func find(n int, a int, b int) int {
	cnt := 0

	if (a > n/2 && b <= n/2) || (a <= n/2 && b > n/2) {
		for n > 1 {
			cnt = cnt + 1
			n /= 2
		}
		return cnt
	} else {
		return find(n/2, a, b)
	}
}

func solution(n int, a int, b int) int {
	return find(n, a, b)
}

func main() {
	fmt.Println(solution(100, 21, 88))
	fmt.Println(solution(2, 1, 2))
	fmt.Println(solution(8, 3, 7))
}
