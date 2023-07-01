package main

import "fmt"

type Node struct {
	data int
	char rune
	val  float64
}

func main() {
	Node1 := Node{data: 1, char: 'a', val: 1.1}

	fmt.Println("Node1.data =", Node1.data)
	fmt.Printf("Node1.char = %c\n", Node1.char)
	fmt.Println("Node1.val =", Node1.val)
}
