package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allposts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "song"},
		Post{Id: 2, Content: "Bonjour Monde!", Author: "dong"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "myung"},
		Post{Id: 4, Content: "Greetings!", Author: "jjang"},
	}

	writer := csv.NewWriter(csvFile)

	for _, post := range allposts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	//레코드 내에 모든 필드가 없더라도 문제가 안되기 위함
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
