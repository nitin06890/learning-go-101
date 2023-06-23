package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Println("Enter a number between 1 to 500: ")
	fmt.Scanf("%d", &num)

	switch {
	case num >= 1 && num <= 100:
		fmt.Println("The number is between 1 and 100")
	case num >= 101 && num <= 200:
		fmt.Println("The number is between 101 and 200")
	case num >= 201 && num <= 300:
		fmt.Println("The number is between 201 and 300")
	case num >= 301 && num <= 400:
		fmt.Println("The number is between 301 and 400")
	case num >= 401 && num <= 500:
		fmt.Println("The number is between 401 and 500")
	default:
		fmt.Println("The number is not between 1 and 500")
	}
}
