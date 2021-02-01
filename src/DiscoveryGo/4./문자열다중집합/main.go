package main

import (
	"fmt"
	"strings"
)

type MultiSet map[string]int

func (m MultiSet) Insert(val string) {
	m[val]++
}

func (m MultiSet) Erase(val string) {
	if m[val] <= 1 {
		delete(m, val)
	} else {
		m[val]--
	}
}

func (m MultiSet) Count(val string) int {
	return m[val]
}

func (m MultiSet) String() string {
	s := "{"
	for val, count := range m {
		s += strings.Repeat(val+" ", count)
	}
	return s + "}"
}

func main() {
	m := MultiSet{}
	m.Insert("song")
	m.Insert("song")
	m.Insert("song")
	fmt.Println(m)
	fmt.Println(m.Count("song"))
}
