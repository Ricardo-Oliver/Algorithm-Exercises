/*21) Faça um algoritmo que leia um determinado ano e mostre se ele é ou não
BISSEXTO.*/

package main

import "fmt"

func main() {
	var year int

	fmt.Print("Enter the year: ")
	fmt.Scan(&year)

	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Printf("The year %v is a leap year.", year)
	} else {
		fmt.Printf("The year %v is not a leap year.", year)
	}
}
