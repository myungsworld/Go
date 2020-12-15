package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	h := &dataStruct.Heap{}

	// [-1,3,-1,5,4], 2번째 큰값
	nums := []int{-1, 3, -1, 5, 4}
	for i := 0; i < 5; i++ {
		h.Push(nums[i])
		if h.Count() > 2 {
			h.Pop()
		}
	}
	fmt.Println(h.Pop())

	// [2,4,-2,-3,8], 제일 큰값
	h = &dataStruct.Heap{}

	nums = []int{2, 4, -2, -3, 8}
	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 1 {
			h.Pop()
		}
	}
	fmt.Println(h.Pop())

	//[-5,-3,1] , 3번째 큰값
	h = &dataStruct.Heap{}
	nums = []int{-5, -3, 1}
	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
	}
	if h.Count() > 3 {
		h.Pop()
	}
	fmt.Println(h.Pop())
}
