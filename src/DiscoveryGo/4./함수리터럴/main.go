package main

import "fmt"

func main() {
	printhello := func() {
		fmt.Println("Hello World")
	}
	func() {
		fmt.Println("Hello 2world")
	}()
	printhello()
}
