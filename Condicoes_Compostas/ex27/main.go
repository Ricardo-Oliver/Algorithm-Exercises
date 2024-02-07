/*27) Crie um programa que leia duas notas de um aluno e calcule a sua média,
mostrando uma mensagem no final, de acordo com a média atingida:
- Média até 4.9: REPROVADO
- Média entre 5.0 e 6.9: RECUPERAÇÃO
- Média 7.0 ou superior: APROVADO.*/

package main

import "fmt"

func main() {
	var (
		note1 float64
		note2 float64
		media float64
	)

	fmt.Print("Enter the nota 1: ")
	fmt.Scan(&note1)
	fmt.Print("Enter the nota 2: ")
	fmt.Scan(&note2)

	media = (note1 + note2) / 2

	if media <= 4.9 {
		fmt.Print("Failed student.")
	} else if media >= 5.0 && media <= 6.9 {
		fmt.Print("Recovery student.")
	} else {
		fmt.Print("Approved student!")
	}
}
