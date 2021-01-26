package main

import "fmt"

var (
	start = rune(44032)
	end   = rune(55204)
)

func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false
	for _, r := range s {
		if start <= r && r < end {
			index := int(r - start)
			result = index%numEnds != 0
		}
	}
	return result
}

func main() {
	var fruits = [...]string{"사과", "바나나", "토마토", "귤"}

	for _, fruit := range fruits {
		if HasConsonantSuffix(fruit) {
			fmt.Printf("%s은 맛있다\n", fruit)
		} else {
			fmt.Printf("%s는 맛있다\n", fruit)
		}
	}
}
