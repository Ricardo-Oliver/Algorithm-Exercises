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
