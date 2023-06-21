package main

import "fmt"

func main() {
	var seconds, hr, min, sec int
	var err error

	fmt.Println("Please enter the number of seconds: ")
	_, err = fmt.Scanf("%d", &seconds)
	if err != nil {
		fmt.Println(err)
		return
	}

	hr = seconds / 3600
	sec = seconds % 3600

	min = sec / 60
	sec = sec % 60

	fmt.Printf("%d seconds = %d hours, %d minutes and %d seconds\n", seconds, hr, min, sec)

}