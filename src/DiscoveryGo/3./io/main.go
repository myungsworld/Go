package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("songdongmyung")

	buffer := make([]byte, 10)

	n, err := io.ReadFull(reader, buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("버퍼 파이트 수: %d\n", n)
	fmt.Println(string(buffer))
}
