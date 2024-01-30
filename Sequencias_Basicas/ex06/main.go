/*6) Faça um programa que leia um número inteiro e mostre o seu antecessor e seu
sucessor.
Ex:
Digite um número: 9
O antecessor de 9 é 8
O sucessor de 9 é 10*/

package main

import "fmt"

func main() {
	var number int
	var predecessor int
	var successor int

	fmt.Print("Enter a value: ")
	fmt.Scan(&number)

	predecessor = number - 1
	successor = number + 1

	fmt.Printf("The predecessor of %v is %v\n", number, predecessor)
	fmt.Printf("The successor of %v is %v", number, successor)
}
