package main

import (
	"fmt"
	"sort"
)

func count(s string, codeCount map[rune]int) {
	for _, r := range s {
		codeCount[r]++
	}
}

func main() {

	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	var keys sort.IntSlice
	for key := range codeCount {
		fmt.Println(string(key))

		keys = append(keys, int(key))
	}
	fmt.Println(keys)

}
