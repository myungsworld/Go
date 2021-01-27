package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteTo(w io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}

func ReadFrom(r io.Reader, lines *[]string) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
func main() {

	f, err := os.Open("abc.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var s string
	fmt.Fscanf(f, "%s\n", &s)
}
