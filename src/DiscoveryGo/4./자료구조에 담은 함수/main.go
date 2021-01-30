package main

import (
	"fmt"
	"strconv"
	"strings"
)

type BinOp func(int, int) int
type StrSet map[string]struct{}
type PrecMap map[string]StrSet

func NewEvaluator(opMap map[string]BinOp, prec PrecMap) func(expr string) int {
	return func(expr string) int {
		return Eval(opMap, prec, expr)
	}
}

func NewStrSet(strs ...string) StrSet {
	m := StrSet{}
	for _, str := range strs {
		m[str] = struct{}{}

	}
	return m
}

func Eval(opMap map[string]BinOp, prec PrecMap, expr string) int {
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
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				return
			}
		b, a := pop(), pop()
		if f := opMap[op]; f != nil {
			nums = append(nums, f(a, b))
		}
		

		for _, token := range strings.Split(expr, " ") {
			if token == "(" {
				ops = append(ops, token)
			} else if _, ok := prec[token]; ok {
				reduce(token)
				ops = append(ops, token)
			} else if token == ")" {
				reduce(token)
			} else {
				num, _ := strconv.Atoi(token)
				nums = append(nums, num)
			}
		}
	reduce(")")
	return nums[0]
	}

	// ... 생략 ...
}

func main() {

	fmt.Println(NewStrSet("Abc", "bdc", "kkk"))
	opMap := map[string]BinOp{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
}
