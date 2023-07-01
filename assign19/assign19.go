package main

import "fmt"

type Node struct {
	data int
}

func main() {
	var NodePtr *[5]Node

	NodePtr = new([5]Node)
	if NodePtr == nil {
		fmt.Println("Memory allocation failed")
		return
	} else {
		fmt.Println("Memory allocation succeeded")
	}
}
