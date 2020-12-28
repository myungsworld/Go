package main

import (
	"net/http"

	"github.com/todolistW/myapp"
)

func main() {
	h := myapp.MakeNewHandler()
	http.ListenAndServe(":3000", h)
}
