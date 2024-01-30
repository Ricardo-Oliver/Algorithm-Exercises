/*16) [DESAFIO] Escreva um programa para calcular a redução do tempo de vida de um
fumante. Pergunte a quantidade de cigarros fumados por dias e quantos anos ele
já fumou. Considere que um fumante perde 10 min de vida a cada cigarro. Calcule
quantos dias de vida um fumante perderá e exiba o total em dias.*/

package main

import "fmt"

func main() {
	var (
		cigarrettesPerDay int
		numberCigarrettes int
		yearsSmoker       int
		lostLifeTime      int
		lostDay           int
	)

	const minutesPerDay int = 1440

	fmt.Print("Enter the number of cigarrettes smoked per day: ")
	fmt.Scan(&cigarrettesPerDay)

	fmt.Print("Tell us how many years you have been smoker: ")
	fmt.Scan(&yearsSmoker)

	numberCigarrettes = (yearsSmoker * 365) * cigarrettesPerDay

	lostLifeTime = numberCigarrettes * 10

	lostDay = lostLifeTime / minutesPerDay

	fmt.Printf("The amount of lost day smoker is %v", lostDay)
}
