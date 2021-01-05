package main

import "fmt"

func isPrimeNum(n int) int {

	last := n / 2
	if n <= 1 {
		return 0
	}

	for i := 2; i <= last; i++ {
		if (n % i) == 0 {
			return 0
		}
	}
	return 1
}

func solution(nums []int) int {
	cnt := 0
	sum := 0
	sum2 := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for j := 1; j < len(nums); j++ {
			if i != j {
				sum += nums[j]
				sum2 = sum
			} else {
				break
			}
			for k := 2; k < len(nums); k++ {
				if i != j && j != k && i != k {
					sum += nums[k]
					fmt.Println(sum)
					cnt += isPrimeNum(sum)
					sum = sum2

				}
			}
		}

	}

	return cnt
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)
	fmt.Println(solution(nums))
}
