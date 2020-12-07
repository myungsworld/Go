package main

import "fmt"

func main() {

	// s := "Hello 월드"

	// s2 := []rune(s)
	// for i := 0; i < len(s2); i++ {
	// 	fmt.Print(string(s2[i]), ",")
	// }

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(arr))

	for i := 0; i < len(arr)/2; i++ {

		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
		fmt.Println(arr)
	}

}
