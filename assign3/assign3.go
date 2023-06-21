package main

import "fmt"

func main() {
	var fName, lName string
	var err error

	fmt.Println("Enter your first name: ")
	_, err = fmt.Scanf("%s", &fName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Enter your last name: ")
	_, err = fmt.Scanf("%s", &lName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Hello %s %s !\n", fName, lName)
}
