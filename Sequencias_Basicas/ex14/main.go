/*14) A locadora de carros precisa da sua ajuda para cobrar seus serviços. Escreva
um programa que pergunte a quantidade de Km percorridos por um carro alugado e a
quantidade de dias pelos quais ele foi alugado. Calcule o preço total a pagar,
sabendo que o carro custa R$90 por dia e R$0,20 por Km rodado.*/

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
