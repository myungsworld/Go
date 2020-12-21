package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreateAt  time.Time `json:"created_at"`
}

type fooHandler struct{}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//writer에 문자열을 출력하는 것
	fmt.Fprint(w, "Hello World")
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	//user의 값을 decode 했고 err가 null인 경우 밑으로 내려와서 수행
	user.CreateAt = time.Now()

	data, _ := json.Marshal(user)
	fmt.Println("1")
	fmt.Println(data)
	fmt.Println("2")
	fmt.Println(string(data))
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	//data는 byte array이기때문에 String으로 바꿔줌
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

//http.Handler 는 serverhttp 함수를 가지고 있는 interface
func NewHttpHandler() http.Handler {
	//router
	mux := http.NewServeMux()

	//func 직접 등록
	mux.HandleFunc("/", IndexHandler)

	mux.HandleFunc("/bar", barHandler)

	//instance 형태 등록
	mux.Handle("/foo", &fooHandler{})
	return mux
}
