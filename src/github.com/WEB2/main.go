package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	//첫번째 인자 id
	uploadFile, header, err := r.FormFile("upload_file")
	//fmt.Println("FormFile의 uploadFIle , header ", uploadFile, header)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()

	//저장할 폴더
	dirname := "./uploads"
	os.MkdirAll(dirname, 0777)
	//Sprintf : "..." 안의 값을 리턴
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	//파일만들기
	file, err := os.Create(filepath)
	//만들어진 파일은 항상 닫아줘야함
	defer file.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	// 뒤의 uploadFile을 file에 넣음
	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}

//static web server (File upload)
func main() {
	//
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.HandleFunc("/uploads", uploadsHandler)

	http.ListenAndServe(":3000", nil)
}
