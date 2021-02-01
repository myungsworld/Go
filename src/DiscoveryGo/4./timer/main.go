package main

import (
	"fmt"
	"time"
)

func CountDown(seconds int) {
	for seconds > 0 {
		fmt.Println(seconds)
		time.Sleep(time.Second)
		seconds--
	}

}

func main() {
	// timer :=
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Let's Begin!")
	})
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("I'm so excited!!")
		// timer.Stop()
	})
	fmt.Println("Ladies and gentlemen!")
	CountDown(5)

}
