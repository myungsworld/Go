package main

import "fmt"

func solution(s string) string {
	if len(s)%2 == 0 {
		even := s[len(s)/2-1 : len(s)/2+1]
		return even
	} else {
		odd := string(s[len(s)/2])
		return odd
	}

}

func main() {
	fmt.Println(solution("abcde"))
	fmt.Println(solution("qwer"))
}
