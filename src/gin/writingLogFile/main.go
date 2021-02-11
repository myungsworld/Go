package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()

	// log를 gin.log에 출력하도록 설정
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Run()
}
