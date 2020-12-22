package main

import (
	"net/http"

	"github.com/WEB3/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
