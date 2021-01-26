package main

import "fmt"

func Example_append() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...)
	fmt.Println(f3)
	// Output:
	// [사과 바나나 토마토 포도 딸기]
}

func Example_sliceCopy() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))
	for i := range src {
		dest[i] = src[i]
	}
	fmt.Println(dest)
	// Output:
	// .

	dest1 := make([]int, len(src))
	copy(dest1, src)
}
