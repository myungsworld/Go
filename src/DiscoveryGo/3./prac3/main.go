package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}

	// bufferedReader := bufio.NewReader(f)
	// byteSlice := make([]byte, 1048576)
	// bufferedReader.Read(byteSlice)
	// fmt.Printf("Read bytes : %s\n", byteSlice)

	//한줄씩 읽는방법
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Printf("Read Line: %s\n", scanner.Text())
	}
	f.Close()
}
