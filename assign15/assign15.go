package main

import "fmt"

func main() {
	var num int = 3
	var numPtr *int = &num

	fmt.Printf("The value of int variable is %d.\n", num)
	fmt.Printf("The value of the pointer to the int variable is %p.\n", numPtr)
	fmt.Printf("The memory address of the int variable is %p.\n", &num)
	fmt.Printf("The value held at the memory location that the pointer is pointing to is %d.\n", *numPtr)
}
