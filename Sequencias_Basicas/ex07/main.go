package main

import "fmt"

func main() {
	var number float64
	var double float64
	var third float64

	fmt.Print("Enter the number: ")
	fmt.Scan(&number)

	double = number * 2
	third = number / 3
	fmt.Printf("The double of %.1f is %.1f\n", number, double)
	fmt.Printf("The third part of %.1f is %.5f", number, third)
}
