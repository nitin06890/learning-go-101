package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("testfile.txt", os.O_WRONLY|os.O_CREATE, 0700)
	if err != nil {
		fmt.Println("Open operation failed")
		os.Exit(1)
	} else {
		fmt.Println("Open operation succeeded")
	}
	_ = file
}
