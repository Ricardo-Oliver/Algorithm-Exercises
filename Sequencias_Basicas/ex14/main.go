package main

import "fmt"

func main() {
	var (
		kmTraveled int
		daysRented int
	)

	var totalPrice float64

	fmt.Print("Enter the number of km traveled: ")
	fmt.Scan(&kmTraveled)

	fmt.Print("Enter the number of days rented: ")
	fmt.Scan(&daysRented)

	totalPrice = (float64(kmTraveled) * 0.20) + (float64(daysRented) * 90)

	fmt.Printf("The total amount to be paid for the car rental is U$%.2f", totalPrice)
}
