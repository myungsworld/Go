package main

import (
	"fmt"
)

func main() {
	var i int
	j := 0
	for i = 1; i < 10; i++ {

		//5단만 없애라
		// if i == 5 {
		// 	continue
		// }

		fmt.Printf("\n%d단\n", i)
		for j = 1; j < 10; j++ {

			//각 단에서 곱하기 5만 없애라
			// if j == 5 {
			// 	continue
			// }

			// 3 x 2 만 없애라
			// if i == 3 && j == 2 {
			// 	continue
			// }

			fmt.Printf("%d x %d = %d", i, j, i*j)
			fmt.Println()
		}
	}
}
