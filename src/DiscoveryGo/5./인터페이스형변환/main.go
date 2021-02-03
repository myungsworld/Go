package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Join(sep string, a ...interface{}) string {
	if len(a) == 0 {
		return ""
	}
	t := make([]string, len(a))
	for i := range a {
		switch x := a[i].(type) {
		case string:
			t[i] = x
		case int:
			t[i] = strconv.Itoa(x)
		case fmt.Stringer:
			t[i] = x.String()
		}
	}
	return strings.Join(t, sep)
}

func main() {
	sep := "abcd"
	fmt.Println(Join(sep, "efc", 1, "사랑"))
}
