package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/urfave/negroni"

	"github.com/unrolled/render"

	"github.com/gorilla/pat"
)

var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "song", Email: "myungsworld@gmail.com"}

	rd.JSON(w, http.StatusOK, user)

	// 위 한줄 코드로 구현 가능
	// w.Header().Add("content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))
}

func addUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		rd.Text(w, http.StatusBadRequest, err.Error())
		// 같음
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprint(w, err)
	}
	user.CreatedAt = time.Now()

	rd.JSON(w, http.StatusOK, user)

	// 위 한줄 코드로 구현 가능
	// w.Header().Add("content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//tmpl, err := template.New("Hello").ParseFiles("templates/hello.tmpl")
	// if err != nil {
	// 	rd.Text(w, http.StatusInternalServerError, err.Error())
	// 	// w.WriteHeader(http.StatusBadRequest)
	// 	// fmt.Fprint(w, err)
	// }
	user := User{Name: "song", Email: "myungsworld@gmail.com"}

	rd.HTML(w, http.StatusOK, "body", user)
	//tmpl.ExecuteTemplate(w, "hello.tmpl", "song")
}

func main() {
	rd = render.New(render.Options{
		// 폴더명을 바꾸고 싶을때
		// Directory:  "template",

		Extensions: []string{".html", ".tmpl"},
		Layout:     "hello",
	})
	mux := pat.New()
	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserInfoHandler)
	mux.Get("/hello", helloHandler)
	n := negroni.Classic()
	n.UseHandler(mux)
	//negroni로 인해 필요없어짐
	// mux.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":3000", n)
}
