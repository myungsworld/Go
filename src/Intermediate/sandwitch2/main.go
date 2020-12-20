package main

import "fmt"

type SpoonOfJam interface {
	//입력    출력
	String() string
}

type Jam interface {
	//입력         출력
	GetOneSpoon() SpoonOfJam
}

type Bread struct {
	val string
}

type StrawberryJam struct {
}

type OrangeJam struct {
}

type AppleJam struct {
}

type SpoonOfStrawberryJam struct {
}

type SpoonOfOrangeJam struct {
}

type SpoonOfAppleJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ Strawberry"
}

func (s *SpoonOfOrangeJam) String() string {
	return "+ Orange"
}

func (s *SpoonOfAppleJam) String() string {
	return "+ Apple"
}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

func (a *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread " + b.val
}

func main() {
	bread := &Bread{}
	//jam := &OrangeJam{}
	jam := &AppleJam{}
	bread.PutJam(jam)
	fmt.Println(bread)
}
