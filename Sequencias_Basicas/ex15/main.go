/*15) Crie um programa que leia o número de dias trabalhados em um mês e mostre o
salário de um funcionário, sabendo que ele trabalha 8 horas por dia e ganha R$25
por hora trabalhada.*/

package main

import "fmt"

func main() {
	var daysWorked int
	var salary int

	fmt.Print("Enter the number of days worked by the employee: ")
	fmt.Scan(&daysWorked)

	salary = (8 * 25) * daysWorked

	fmt.Printf("The salary employee is U$%v", salary)
}
