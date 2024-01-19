package main

import "fmt"

func main() {
	var name string 
	var salary float64

	fmt.Println("What is your name ?")
	fmt.Scan(&name)
	fmt.Println("What is your salary ?")
	fmt.Scan(&salary)

	fmt.Printf("Employee name: %v\n", name)
	fmt.Printf("Salary: %v\n", salary)
	fmt.Printf("Employee %v has a salary of R$%.2f in June", name, salary)
}
