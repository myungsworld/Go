package main

import "fmt"

type Bread struct {
	val string
}

type Jam struct {
}

//이거 추가하면 {} 없이 출력
func (b *Bread) String() string {
	return b.val
}

func (b *Bread) PutJam(jam *Jam) {
	b.val += jam.GetVal()
}

func (j *Jam) GetVal() string {
	return " + jam"
}

func main() {
	b := &Bread{val: "bread"}
	j := &Jam{}
	b.PutJam(j)
	fmt.Println(b)
}
