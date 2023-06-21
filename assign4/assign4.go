package main

import "fmt"

func main() {
	const pi float64 = 3.14
	var radius float64
	var err error

	fmt.Println("Enter the radius of the circle: ")
	_, err = fmt.Scanf("%f", &radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The area of the circle is =  %f\n", pi*radius*radius)
}
