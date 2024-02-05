/*22) Escreva um programa que leia o ano de nascimento de um rapaz e mostre a sua
situação em relação ao alistamento militar.
- Se estiver antes dos 18 anos, mostre em quantos anos faltam para o
alistamento.
- Se já tiver depois dos 18 anos, mostre quantos anos já se passaram do
alistamento.*/

package main

import (
	"fmt"
	"time"
)

func main() {
	var yearBirth int
	const enlistment = 18

	fmt.Print("Enter your year of birth: ")
	fmt.Scan(&yearBirth)

	age := time.Now().Year() - yearBirth

	if age < enlistment {
		missingEnlistment := enlistment - age
		fmt.Printf("There is still %d year's left to enlist in military service.", missingEnlistment)
	} else if age > enlistment {
		pastEnlistment := age - enlistment
		fmt.Printf("%d year's have passed since military service enlistment.", pastEnlistment)
	} else {
		fmt.Println("You're your military service enlistment year.")
	}
}
