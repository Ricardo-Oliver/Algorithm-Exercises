package main

import "fmt"

func main() {
	var (
		num1, num2 int
	)

	fmt.Println("Enter a value: ")
	fmt.Scan(&num1)
	fmt.Println("Enter another value: ")
	fmt.Scan(&num2)

	sum := num1 + num2 

	fmt.Printf("The sum between %v and %v is equal to %v", num1, num2, sum)
}