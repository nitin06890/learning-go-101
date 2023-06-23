package main

import "fmt"

func main() {
	var numerator, denominator int
	var err error

	fmt.Println("Please enter a numerator: ")
	_, err = fmt.Scanf("%d", &numerator)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Please enter a denominator: ")
	_, err = fmt.Scanf("%d", &denominator)
	if err != nil {
		fmt.Println(err)
		return
	}
	remainder := numerator % denominator

	if remainder == 0 {
		fmt.Println("There is no remainder")
	} else {
		fmt.Printf("The remainder is %d\n", remainder)
	}
}
