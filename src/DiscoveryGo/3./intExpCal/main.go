package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Eval(expr string) int {

	// ops는 연산자 nums는 정수값
	var ops []string
	var nums []int
	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	reduce := func(higher string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if strings.Index(higher, op) < 0 {
				//목록에 없는 연산자 이므로 종료
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				return
			}
			b, a := pop(), pop()
			switch op {
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, a-b)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, a/b)
			}
		}
	}

	for _, token := range strings.Split(expr, " ") {
		fmt.Println(token)
		switch token {
		case "(":
			ops = append(ops, token)
		case "+", "-":
			//덧셈과 뺄셈 이상의 우선순위를 가진 사칙연산 적용
			reduce("+-*/")
			ops = append(ops, token)
		case "*", "/":
			//곱셉과 나눗셈 이상의 우선순위를 가진 것은 둘뿐
			reduce("*/")
			ops = append(ops, token)
		case ")":
			reduce("+-*?/(")
		default:
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}
	reduce("+-*/")
	return nums[0]
}

func main() {
	fmt.Println(Eval("3 * ( 3 + 1 * 3 ) / 2"))
}
