/*4) Desenvolva um algoritmo que leia dois números inteiros e mostre o somatório
entre eles.
Ex:
Digite um valor: 8
Digite outro valor: 5
A soma entre 8 e 5 é igual a 13.*/

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