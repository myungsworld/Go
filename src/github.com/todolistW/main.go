package main

import (
	"net/http"

	"github.com/todolistW/myapp"
)

func main() {
	m := myapp.MakeNewHandler("./test.db")
	defer m.Close()

	err := http.ListenAndServe(":3000", m)
	if err != nil {
		panic(err)
	}
}
