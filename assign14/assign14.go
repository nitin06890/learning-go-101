package main

import "fmt"

func main() {
	var num int = 3
	var numPtr *int = &num

	fmt.Printf("The value of our pointer is %p.\n", numPtr)
}
