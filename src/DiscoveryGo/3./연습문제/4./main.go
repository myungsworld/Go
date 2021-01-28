package main

import "fmt"

type Queue struct {
	arr []string
}

func (q *Queue) push(s string) {
	q.arr = append(q.arr, s)
}

func main() {
	q := &Queue{}
	q.push("song")
	q.push("dong")
	fmt.Println(q)
}
