package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	path := "C:/Users/myungsworld/go/src/go/src/github.com/WEB2/cmdl.txt"
	//파일 오픈
	file, _ := os.Open(path)
	defer file.Close()

	os.RemoveAll("./uploads")

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	//
	//Base() 는 파일경로에서 이름만 짤라냄 ex) cmdl.txt
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)

	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())
	uploadsHandler(res, req)

	assert.Equal(http.StatusOK, res.Code)

	uploadsFilePath := "./uploads/" + filepath.Base(path)

	_, err = os.Stat(uploadsFilePath)
	assert.NoError(err)

	uploadFile, _ := os.Open(uploadsFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}

	uploadFile.Read(uploadData)
	uploadFile.Read(originData)

	assert.Equal(originData, uploadData)
}
