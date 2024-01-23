package main

import "fmt"

func main() {
	var note1 float64
	var note2 float64
	var average float64

	fmt.Print("Nota 1: ")
	fmt.Scan(&note1)

	fmt.Print("Nota 2: ")
	fmt.Scan(&note2)

	average = (note1 + note2) / 2
	
	fmt.Printf("The average between %.1f and %.1f is equal the %.1f", note1, note2, average)
}