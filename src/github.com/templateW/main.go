package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{Name: "song", Email: "myungsworld@gmail.com", Age: 23}
	user2 := User{Name: "aaa", Email: "aaa@gmail.com", Age: 40}
	users := []User{user, user2}
	//New = tempalte 이름 설정 ParseFiles : tempalte으로 정할 파일 설정 다중으로 가능
	tmpl, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl", "templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)

}
