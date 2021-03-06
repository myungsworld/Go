package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// POST /post?id=1234&page=1 HTTP/1.1
// Content-Type: application/x-www-form-urlencoded

// name=manu&message=this_is_great

func main() {
	r := gin.Default()

	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		email := c.PostForm("email")

		fmt.Printf("id : %s, page : %s, name : %s, email : %s", id, page, name, email)
	})
	r.Run()
}
