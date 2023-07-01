package main

import "fmt"

type Employee struct {
	fName     string
	lName     string
	empId     int
	ninNumber int
	title     string
}

type Employees struct {
	emp1 Employee
	emp2 Employee
}

func main() {
	emp := Employees{}

	fmt.Println("Enter first name of employee 1:")
	fmt.Scanln(&emp.emp1.fName)
	fmt.Println("Enter last name of employee 1:")
	fmt.Scanln(&emp.emp1.lName)
	fmt.Println("Enter employee ID of employee 1:")
	fmt.Scanln(&emp.emp1.empId)
	fmt.Println("Enter the last 4 digits of national insurance number of employee 1:")
	fmt.Scanln(&emp.emp1.ninNumber)
	fmt.Println("Enter title of employee 1:")
	fmt.Scanln(&emp.emp1.title)

	fmt.Println("Enter first name of employee 2:")
	fmt.Scanln(&emp.emp2.fName)
	fmt.Println("Enter last name of employee 2:")
	fmt.Scanln(&emp.emp2.lName)
	fmt.Println("Enter employee ID of employee 2:")
	fmt.Scanln(&emp.emp2.empId)
	fmt.Println("Enter the last 4 digits of national insurance number of employee 2:")
	fmt.Scanln(&emp.emp2.ninNumber)
	fmt.Println("Enter title of employee 2:")
	fmt.Scanln(&emp.emp2.title)

	fmt.Printf("Employee information of %s %s: ID: %d, NIN: %d, Title: %s\n", emp.emp1.fName, emp.emp1.lName, emp.emp1.empId, emp.emp1.ninNumber, emp.emp1.title)
	fmt.Printf("Employee information of %s %s: ID: %d, NIN: %d, Title: %s\n", emp.emp2.fName, emp.emp2.lName, emp.emp2.empId, emp.emp2.ninNumber, emp.emp2.title)

}
