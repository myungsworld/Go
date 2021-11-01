package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("/Users/myungsworld/go/src/test/data/**/*.csv")
	if err != nil {
		panic(err)
	}
	for _ , file := range files {

		f, _ := os.Open(file)
		rdr := csv.NewReader(bufio.NewReader(f))

		rows, _ := rdr.ReadAll()

		values := make(map[string]int)
		for _ , row := range rows {
			values[row[0]] ++
			if values[row[0]] == 2 {

				fmt.Println(row)
				break
			}
		}

	}
}
