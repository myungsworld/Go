package myapp

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/todolistW/model"
	"github.com/urfave/negroni"

	"github.com/unrolled/render"

	"github.com/gorilla/mux"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var rd *render.Render = render.New()

type AppHandler struct {
	http.Handler
	db model.DBHandler
}

//signin에서 session이 저장되면 indexhandler가 호출될때 SessionID를 읽어와야함
func getSessionID(r *http.Request) string {
	session, err := store.Get(r, "session")
	if err != nil {
		return ""
	}
	val := session.Values["id"]
	if val == nil {
		return ""
	}
	return val.(string)
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	// list := []*Todo{}
	// for _, v := range todoMap {
	// 	list = append(list, v)
	// }
	list := a.db.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	// id := len(todoMap) + 1
	// todo := &Todo{id, name, false, time.Now()}
	// todoMap[id] = todo
	todo := a.db.AddTodo(name)
	rd.JSON(w, http.StatusCreated, todo)
}

type Sucess struct {
	Delete bool `json:"delete"`
}

func (a *AppHandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.db.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Sucess{true})
	} else {
		rd.JSON(w, http.StatusOK, Sucess{false})
	}
	// if _, ok := todoMap[id]; ok {
	// 	delete(todoMap, id)
	// 	rd.JSON(w, http.StatusOK, Sucess{true})
	// } else {
	// 	rd.JSON(w, http.StatusOK, Sucess{false})
	// }
}

func (a *AppHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	ok := a.db.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Sucess{true})
	} else {
		rd.JSON(w, http.StatusOK, Sucess{false})
	}
	// if todo, ok := todoMap[id]; ok {
	// 	todo.Completed = complete
	// 	rd.JSON(w, http.StatusOK, Sucess{true})
	// } else {
	// 	rd.JSON(w, http.StatusOK, Sucess{false})
	// }

}

func (a *AppHandler) Close() {
	a.db.Close()
}

func CheckSignin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	//유저가 처음에 로그인을 하려고 한다면
	if strings.Contains(r.URL.Path, "/signin") ||
		strings.Contains(r.URL.Path, "/auth") {
		next(w, r)
		return
	}

	//유저가 로그인을 했다면
	sessionID := getSessionID(r)
	if sessionID != "" {
		next(w, r)
		return
	}

	//유저가 로그인을 하지 않았다면
	http.Redirect(w, r, "/signin.html", http.StatusTemporaryRedirect)

}

func MakeNewHandler(filepath string) *AppHandler {

	r := mux.NewRouter()

	// session에서 가져온 id가 없으면 로그인 화면
	// id가 있으면 다음 => decoreate pattern
	n := negroni.New(NewRecovery(), NewLogger(), negroni.HandlerFunc(CheckSignin), NewStatic(http.Dir("public")))
	n.UseHandler(r)
	a := &AppHandler{
		Handler: n,
		db:      model.NewDBHandler(filepath),
	}
	r.HandleFunc("/", a.indexHandler)
	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")
	r.HandleFunc("/auth/google/login", googleLoginHandler)
	r.HandleFunc("/auth/google/callback", googleAuthCallback)

	return a
}
