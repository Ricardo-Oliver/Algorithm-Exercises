package main

import "fmt"

func main() {
	var daysWorked int
	var salary int

	fmt.Print("Enter the number of days worked by the employee: ")
	fmt.Scan(&daysWorked)

	salary = (8 * 25) * daysWorked

	fmt.Printf("The salary employee is U$%v", salary)
}
