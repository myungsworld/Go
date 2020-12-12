package main

import "fmt"

type Student struct {
	name    string
	class   int
	sungjuk Sungjuk
}

type Sungjuk struct {
	name  string
	grade string
}

func (s *Student) ViewSungjuk() {
	fmt.Println(s.sungjuk)
}

// 위와 똑같은 메서드
func ViewSungjuk(s Student) {
	fmt.Println(s.sungjuk)
}

func (p *Student) InputSungjuk(name string, grade string) {
	p.sungjuk.name = name
	p.sungjuk.grade = grade
}

//위와 똑같은 메서드
func InputSungjuk(p *Student, name string, grade string) {
	p.sungjuk.name = name
	p.sungjuk.grade = grade
}

func main() {
	var s Student
	s.name = "동명"
	s.class = 1

	s.sungjuk.name = "수학"
	s.sungjuk.grade = "C"
	s.ViewSungjuk()
	s.InputSungjuk("과학", "A")
	s.ViewSungjuk()
}
