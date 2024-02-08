/*29) Desenvolva um programa que leia o nome de um funcionário, seu salário,
quantos anos ele trabalha na empresa e mostre seu novo salário, reajustado de
acordo com a tabela a seguir:
- Até 3 anos de empresa: aumento de 3%
- entre 3 e 10 anos: aumento de 12.5%
- 10 anos ou mais: aumento de 20%*/

package main

import "fmt"

func main() {
	var name string
	var salary float64
	var totalWorkedTime int

	fmt.Print("Enter the employee name: ")
	fmt.Scan(&name)
	fmt.Print("Enter the employee salary: ")
	fmt.Scan(&salary)
	fmt.Print("Enter the amount of time in years that the employee has worked at the company: ")
	fmt.Scan(&totalWorkedTime)

	if totalWorkedTime < 3 {
		newSalary := salary + (salary * 0.03)
		fmt.Printf("The new salary of the employee is U$%.2f", newSalary)
	} else if totalWorkedTime >= 3 && totalWorkedTime < 10 {
		newSalary := salary + (salary * 0.125)
		fmt.Printf("The new salary of the employee is U$%.2f", newSalary)
	} else {
		newSalary := salary + (salary * 0.2)
		fmt.Printf("The new salary of the employee is U$%.2f", newSalary)
	}
}
