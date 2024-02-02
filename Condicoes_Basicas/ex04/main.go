/*20) Desenvolva um programa que leia um número inteiro e mostre se ele é PAR ou
ÍMPAR.*/

package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter the number: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Printf("The number %v is par", number)
		return
	}

	fmt.Printf("The number %v is impar", number)
}
