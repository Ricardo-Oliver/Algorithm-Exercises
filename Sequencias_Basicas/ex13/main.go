package main

import "fmt"

func main() {
	var salary float64
	var newSalary float64

	fmt.Print("Enter the employee's salary: U$")
	fmt.Scan(&salary)

	newSalary = salary + (salary * 0.15)

	fmt.Printf("The new employee's salary is U$%.2f", newSalary)
}
