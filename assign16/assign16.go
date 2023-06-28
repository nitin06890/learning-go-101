package main

import "fmt"

func main() {
	var arr [10]int

	fmt.Printf("The value of the array is %d.\n", arr)
	fmt.Printf("The value of the array is %p.\n", &arr[0])
}
