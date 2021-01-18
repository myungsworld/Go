package main

import (
	"fmt"
	"sync"
)

type singleton2 struct {
}

var singleInstance *singleton2
var once sync.Once

func getIntance() *singleton2 {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("single 객체 생성중")
				singleInstance = &singleton2{}
			})
	} else {
		fmt.Println("single 객체 이미 생성됨-2")
	}
	return singleInstance
}

func main() {
	// for i := 0; i < 50; i++ {
	// 	go singleton.GetInstance()
	// }
	// fmt.Scanln()

	for i := 0; i < 50; i++ {
		go getIntance()
	}
	fmt.Scanln()
}
