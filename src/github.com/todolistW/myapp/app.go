package myapp

import (
	"net/http"
	"strconv"
	"time"

	"github.com/unrolled/render"

	"github.com/urfave/negroni"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo

var rd *render.Render

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := []*Todo{}
	for _, v := range todoMap {
		list = append(list, v)
	}
	rd.JSON(w, http.StatusOK, list)
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	id := len(todoMap) + 1
	todo := &Todo{id, name, false, time.Now()}
	todoMap[id] = todo
	rd.JSON(w, http.StatusCreated, todo)
}

type Sucess struct {
	Delete bool `json:"delete"`
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Sucess{true})
	} else {
		rd.JSON(w, http.StatusOK, Sucess{false})
	}
}

func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	if todo, ok := todoMap[id]; ok {
		todo.Completed = complete
		rd.JSON(w, http.StatusOK, Sucess{true})
	} else {
		rd.JSON(w, http.StatusOK, Sucess{false})
	}

}

// func testTodos() {
// 	todoMap[1] = &Todo{1, "take a shit", false, time.Now()}
// 	todoMap[2] = &Todo{2, "take a shi", true, time.Now()}
// 	todoMap[3] = &Todo{3, "take a sh", false, time.Now()}
// }

func MakeNewHandler() http.Handler {
	todoMap = make(map[int]*Todo)
	rd = render.New()
	// testTodos()
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", completeTodoHandler).Methods("GET")
	n := negroni.Classic()
	n.UseHandler(r)
	return n
}
