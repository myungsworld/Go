package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/writeFile/dat1", d1, 0644)
	check(err)
	f, err := os.Create("/writeFile/dat2")
	check(err)
	defer f.Close()

	d2 := []byte{115, 11, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)
}
