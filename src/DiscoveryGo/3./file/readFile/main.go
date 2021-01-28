package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	//f.Close()
	f.WriteString("This is Song from Korea")
	f.Close()

	f1, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	data := make([]byte, 4)

	// 6번째부터 읽어옴
	// Seek의 두번째 인자가 0이면 처음부터 1이면 현재부터 2이면 마지막부터

	f1.Seek(8, 0)
	// file로부터 5글자를 읽어옴

	f1.Read(data)
	fmt.Printf("The file data is %s\n", string(data))
	f1.Close()
}
