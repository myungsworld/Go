package main

import (
	"fmt"
	"os/exec"
)

func main() {

	//goExecPath , err := exec.LookPath("go")
	//if err != nil {
	//	fmt.Println("Error : ", err)
	//} else {
	//	fmt.Println("Go Executable: ", goExecPath)
	//}
	//
	//cmd := &exec.Cmd{
	//	Path: "./test.sh",
	//	Stdout: os.Stdout,
	//	Stderr: os.Stdout,
	//}
	//
	//cmd.Start()
	//
	//for i := 1; i <30000; i++ {
	//	fmt.Println(i)
	//}
	//
	//cmd.Wait()

	cmd := exec.Command("./test.sh")

	if output, err := cmd.Output(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Output: %s\n", output)
	}

}
