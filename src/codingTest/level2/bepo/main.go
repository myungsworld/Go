package main

import "fmt"

func solution(progresses []int, speeds []int) []int {
	var newnum []int
	for len(progresses) > 0 {
		cnt := 1
		for progresses[0] < 100 {
			for i := 0; i < len(progresses); i++ {
				progresses[i] += speeds[i]
			}
		}
		for i := 1; i < len(progresses); i++ {

			if progresses[i] < 100 {
                break
            }
			if progresses[i] >= 100 {
				cnt++
			}
		}
		progresses = progresses[cnt:]
		newnum = append(newnum, cnt)
	}
	return newnum
}

func main() {
	progresses := []int{93, 30, 55}
	speeds := []int{1, 30, 5}
	fmt.Println(solution(progresses, speeds))
}
