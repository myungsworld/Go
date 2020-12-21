package main

import (
	"net/http"

	"github.com/WEB/myapp"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
