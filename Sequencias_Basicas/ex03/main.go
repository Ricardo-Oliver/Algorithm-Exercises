/*3) Crie um programa que leia o nome e o salário de um funcionário, mostrando no
final uma mensagem.
Ex:
Nome do Funcionário: Maria do Carmo
Salário: 1850,45
O funcionário Maria do Carmo tem um salário de R$1850,45 em Junho.*/

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
