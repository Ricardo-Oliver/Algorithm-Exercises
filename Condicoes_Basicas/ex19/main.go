/*19) Crie um algoritmo que leia o nome e as duas notas de um aluno, calcule a sua
média e mostre na tela. No final, analise a média e mostre se o aluno teve ou
não um bom aproveitamento (se ficou acima da média 7.0).*/

package main

import "fmt"

func main() {
	var name string
	var note1 float64
	var note2 float64
	var average float64

	fmt.Println("What is the student's name ?")
	fmt.Scan(&name)

	fmt.Print("Enter note 1: ")
	fmt.Scan(&note1)
	fmt.Print("Enter note 2: ")
	fmt.Scan(&note2)

	average = (note1 + note2) / 2

	if average >= 7 {
		fmt.Printf("The %s student average is %.1f - Aluno APROVADO!!!", name, average)
	} else {
		fmt.Printf("The %s student average is %.1f - Aluno REPROVADO!", name, average)
	}
}
