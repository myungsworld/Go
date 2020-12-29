package main

import (
	"net/http"

	"github.com/restfulW/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
