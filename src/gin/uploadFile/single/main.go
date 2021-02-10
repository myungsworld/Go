package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//업로드 되는 파일의 메모리 사이즈 설정
	//현재 여기서는 8MiB의 사이즈 할당
	router.MaxMultipartMemory = 8 << 20

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		//files라는 이름으로 저장됨
		c.SaveUploadedFile(file, "./files")

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run()
}

// curl -X POST http://localhost:8080/upload \
//   -F "file=@/Users/myungsworld/aws/README.md" \
//   -H "Content-Type: multipart/form-data"
