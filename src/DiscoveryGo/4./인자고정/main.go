package main

import (
	"bufio"
	"fmt"
	"io"
)

//같은 원소가 여러번 들어갈수 있는 집합인 MultiSet을 기본적으로 제공하는 맵을 이용하여 만들기

type MultiSet map[string]int
type SetOp func(m MultiSet, val string)

func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func NewMultiSet() map[string]int {
	return map[string]int{}
}

//Insert 함수는 집합에 val을 추가한다.
func Insert(m MultiSet, val string) {
	m[val]++
}

//Erase 함수는 집합에서 val을 제거한다. 집합에 val이 없는 경우에는 아무 일도 일어나지 않는다.
func Erase(m map[string]int, val string) {
	_, ok := m[val]
	if ok {
		delete(m, val)
	}

}

//Count 함수는 집합에 val이 들어 있는 횟수를 구한다.
func Count(m map[string]int, val string) int {
	for key, value := range m {
		if key == val {
			return value
		}
	}
	return 0
}

//String 함수는 집합에 들어 있는 원소들을 {}안에 빈 칸으로 구분하여 넣은 문자열을 반환한다.
func String(m map[string]int) string {
	return ""
}

func main() {
	m := NewMultiSet()

	Insert(m, "3")
	Insert(m, "3")
	Insert(m, "song")
	fmt.Println("3의 갯수는", Count(m, "3"))
	fmt.Println(m)
	fmt.Println("Erase(m.3)실행")
	Erase(m, "3")
	fmt.Println(m)
	fmt.Println("3의 갯수는", Count(m, "3"))
}
