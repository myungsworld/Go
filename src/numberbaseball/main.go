package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	strikes int
	balls   int
}

func main() {
	//난수를 만들때 Seed 값을 시간으로 두면 시간은 계속 변하니까 난수 생성가능
	rand.Seed(time.Now().UnixNano())

	//무작위 숫자 3개
	numbers := MakeNumbers()

	cnt := 0
	for {
		cnt++
		//사용자 입력 받기
		inputNumbers := InputNumbers()

		// 결과 비교
		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)
		//3s 인가?
		if IsGameEnd(result) {
			// 게임끝
			break
		}
	}

	// 게임 끝. 몇번만에 맞췄는지 출력.
	fmt.Printf("%d번만에 맞췄습니다.", cnt)
}

func MakeNumbers() [3]int {
	// 0에서 9사이 안겹치는 무작위 숫자 반환
	var rst [3]int
	for i := 0; i < 3; i++ {
		for {
			n := rand.Intn(10)
			duplicated := false
			for j := 0; j < i; j++ {
				if rst[j] == n {
					duplicated = true
					break
				}
			}
			if !duplicated {
				rst[i] = n
				break
			}
		}

	}
	//fmt.Println(rst)
	return rst
}

func InputNumbers() [3]int {
	// 키보드로부터 0~9 사이의 겹치지 않는 3개 숫자 반환
	var rst [3]int
	for {
		fmt.Println("겹치지 않는 0~9 사이의 숫자 3개를 입력하세요.")
		var num int
		_, err := fmt.Scanf("%d\n", &num)

		if err != nil {
			fmt.Println("잘못 입력 하셨습니다.")
			continue
		}
		success := true
		idx := 0
		for num > 0 {
			n := num % 10
			num = num / 10

			duplicated := false
			for j := 0; j < idx; j++ {
				if rst[j] == n {
					duplicated = true
					break
				}
			}
			if duplicated {
				fmt.Println("중복된 값이 있습니다. 다시 입력하세요.")
				success = false
				break
			}
			if idx >= 3 {
				// fmt.Println("3개 보다 많은 갑을 입력하셨습니다. 다시 입력하세요.")
				success = false
				break
			}
			rst[idx] = n
			idx++
		}
		if idx < 3 {
			// fmt.Println("3개의 숫자를 입력하세요.")
			success = false
		}

		if !success {
			continue
		}
		break
	}
	rst[0], rst[2] = rst[2], rst[0]
	fmt.Println(rst)
	return rst
}

func CompareNumbers(numbers, inputNumbers [3]int) Result {
	// 두개의 숫자 3개 비교해서 결과 반한
	strikes := 0
	balls := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if numbers[i] == inputNumbers[j] {
				if i == j {
					strikes++
				} else {
					balls++
				}
			}
		}
	}

	return Result{strikes, balls}
}

func PrintResult(result Result) {
	fmt.Printf("%dS %dB\n", result.strikes, result.balls)
}

func IsGameEnd(result Result) bool {
	// 비교 결과가 3스트라이크 인지 확인
	return result.strikes == 3
}
