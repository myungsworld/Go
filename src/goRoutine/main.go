package main

import (
	"fmt"
	"sync"
)


var wg sync.WaitGroup

func main() {

	fmt.Println("이것은 고 루틴을 들어가기전")

	wg.Add(2)

	var i int

	go func() {
		defer wg.Done()
		for a := 0 ; a < 1000 ; a++ {
			i++
			fmt.Println("이거슨 첫번쨰")
			fmt.Print(i)
		}
	}()

	go func() {
		defer wg.Done()
		for a := 0 ; a < 1000 ; a++ {
			i++
			fmt.Println("이거슨 두번째")
			fmt.Print(i)
		}
	}()

	wg.Wait()
}
