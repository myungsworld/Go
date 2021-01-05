package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "song"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "dong"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "myung"}
	post4 := Post{Id: 4, Content: "Greetings!", Author: "jjang"}
	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["song"] {
		fmt.Println(post)
	}

	for _, post := range PostsByAuthor["myung"] {
		fmt.Println(post)
	}
}
