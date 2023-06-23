package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, x1, x2 float64

	fmt.Println("Enter the value of variable A to solve the quadratic formula: ")
	_, err := fmt.Scanf("%f", &a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Enter the value of variable B to solve the quadratic formula: ")
	_, err = fmt.Scanf("%f", &b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Enter the value of variable C to solve the quadratic formula: ")
	_, err = fmt.Scanf("%f", &c)
	if err != nil {
		fmt.Println(err)
		return
	}

	x1 = (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	x2 = (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)

	fmt.Printf("The solution using the '+' operator is %f and the soltuion using the '-' operator is %f\n", x1, x2)
}
