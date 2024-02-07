/*18) Faça um programa que leia o ano de nascimento de uma pessoa, calcule a idade
dela e depois mostre se ela pode ou não votar.*/

package main

import (
	"fmt"
	"time"
)

func main() {
	var yearBirth int
	var age int

	fmt.Print("Enter your year of birth: ")
	fmt.Scan(&yearBirth)

	age = time.Now().Year() - yearBirth

	if age >= 16 {
		fmt.Print("You can now vote in the elections")
		return
	} else {
		fmt.Print("You cannot participate in elections yet.")
	}
}
