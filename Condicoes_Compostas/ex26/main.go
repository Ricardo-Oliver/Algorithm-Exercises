/*26) Escreva um algoritmo que leia dois números inteiros e compare-os, mostrando
na tela uma das mensagens abaixo:
- O primeiro valor é o maior
- O segundo valor é o maior
- Não existe valor maior, os dois são iguais*/

package main

import "fmt"

func main() {
	var (
		firstNumber  int
		secondNumber int
	)

	fmt.Print("Enter the first number: ")
	fmt.Scan(&firstNumber)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&secondNumber)

	if firstNumber > secondNumber {
		fmt.Println("The first value is the largest.")
	} else if secondNumber > firstNumber {
		fmt.Println("The second value is the largest.")
	} else {
		fmt.Println("There is no greater value, both are equal.")
	}
}
