package main

import (
	"bytes"
	"fmt"
)

func solution(s string) string {
	var buffer bytes.Buffer
	last4 := s[len(s)-4:]
	for i := 0; i < len(s)-4; i++ {
		buffer.WriteString("*")
	}
	newstring := fmt.Sprintf(buffer.String() + last4)
	return newstring
}

func main() {
	fmt.Println(solution("01020467089"))
	fmt.Println(solution("233414422223333"))
}
