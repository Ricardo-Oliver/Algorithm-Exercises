/*13) Faça um algoritmo que leia o salário de um funcionário, calcule e mostre o
seu novo salário, com 15% de aumento.*/

package main

import "fmt"

func main() {
	var salary float64
	var newSalary float64

	fmt.Print("Enter the employee's salary: U$")
	fmt.Scan(&salary)

	newSalary = salary + (salary * 0.15)

	fmt.Printf("The new employee's salary is U$%.2f", newSalary)
}
