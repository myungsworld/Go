package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func (t *Student) SetName(newname string) {
	t.name = newname
}

func (t *Student) SetAge(newage int) {
	t.age = newage
}

func main() {
	var a *Student
	a = &Student{"aaa", 20, 10}
	// a가 instance
	a.SetName("bbb")
	a.SetAge(11)
	fmt.Println(a)
}
