package main

import (
	"fmt"
	"time"
)

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

type Task struct {
	title  string
	status status
	due    *time.Time
}

type ByteSize float64

const (
	_           = iota // ignore first value
	KB ByteSize = 1 << (10 * iota)
	MB
	GT
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	task := Task{
		title:  "lol",
		status: TODO,
	}

	fmt.Println(task)
	fmt.Println(KB)
}
