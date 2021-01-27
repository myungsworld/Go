package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	s := "Hello World"
	r := strings.NewReader(s)

	w := bufio.NewWriter(f)

	w.ReadFrom(r)
	w.Flush()
}
