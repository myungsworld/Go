package main

import (
	"fmt"
	"os"
)

func main() {
	googleid := os.Getenv("GOOGLE_CLIENT_ID")
	fmt.Print("1")
	fmt.Print(googleid)
}
