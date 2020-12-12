package main

import "fmt"

func main() {
	arr := [11]int{6, 2, 3, 5, 7, 1, 2, 9, 6, 4, 3}
	temp := [10]int{}

	for i := 0; i < len(arr); i++ {
		idx := arr[i]
		temp[idx]++
	}
	fmt.Println(temp)
	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[idx] = i
			idx++
		}
	}
	fmt.Println(arr)
}
