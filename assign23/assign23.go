package main

import "fmt"

type Example struct {
	integer int
}

func main() {
	var ptr *Example
	var test Example
	ptr = &test

	test.integer = 1
	fmt.Println(test.integer)

	(*ptr).integer = 2
	fmt.Println(test.integer)

	ptr.integer = 3
	fmt.Println(test.integer)
}
