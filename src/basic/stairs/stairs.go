package main

import "fmt"

func main() {
	// i := 0
	// j := 0

	// for i = 0; i < 5; i++ {
	// 	for j = 0; j <= i; j++ {
	// 		fmt.Printf("*")
	// 	}
	// 	fmt.Println()
	// }

	// i = 0
	// j = 0

	// for i = 0; i < 5; i++ {
	// 	for j = 0; j < 5-i; j++ {
	// 		fmt.Printf("*")
	// 	}
	// 	fmt.Println()
	// }

	//다이아몬드
	for i := 0; i < 4; i++ {
		for j := 0; j < 3-i; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < 5-2*i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}
