package main

import (
	"net/http"

	"github.com/todolistW/myapp"
	"github.com/urfave/negroni"
)

func main() {
	m := myapp.MakeNewHandler("./test.db")
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	http.ListenAndServe(":3000", n)
}
