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
	Title  string
	Status status
	*Deadline
}

type Deadline time.Time

func (d *Deadline) OverDue() bool {
	return d != nil && time.Time(*d).Before(time.Now())
}

// func (t Task) OverDue() bool {
// 	return t.Deadline.OverDue()
// }

// func ExampleDeadline_OverDue() {
// 	d1 := Deadline(time.Now().Add(-4 * time.Hour))
// 	d2 := Deadline(time.Now().Add(4 * time.Hour))
// 	fmt.Println(d1.OverDue())
// 	fmt.Println(d2.OverDue())
// 	// Output:
// 	// .
// }

func Example_taskTestAll() {
	d1 := Deadline(time.Now().Add(-4 * time.Hour))
	d2 := Deadline(time.Now().Add(4 * time.Hour))
	t1 := Task{"4h ago", TODO, &d1}
	t2 := Task{"4h later", TODO, &d2}
	t3 := Task{"No due", TODO, nil}

	fmt.Println(t1.OverDue())
	fmt.Println(t2.OverDue())
	fmt.Println(t3.OverDue())
	// Output:
	// .
}
