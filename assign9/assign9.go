package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args

	if len(arg) != 3 {
		fmt.Println("The arg slice usage: ./assign9 FirstName LastName")
		return
	}
	fmt.Println("Hello,", arg[1], arg[2])
}