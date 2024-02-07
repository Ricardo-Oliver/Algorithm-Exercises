/*25) [DESAFIO] Crie um programa que leia o tamanho de três segmentos de reta.
Analise seus comprimentos e diga se é possível formar um triângulo com essas
retas. Matematicamente, para três segmentos formarem um triângulo, o comprimento
de cada lado deve ser menor que a soma dos outros dois.*/

package main

import (
	"fmt"
)

func main() {
	var (
		segmentA int
		segmentB int
		segmentC int
	)

	fmt.Print("Enter the segment A: ")
	fmt.Scan(&segmentA)

	fmt.Print("Enter the segment B: ")
	fmt.Scan(&segmentB)

	fmt.Print("Enter the segment C: ")
	fmt.Scan(&segmentC)

	if segmentA < segmentB+segmentC && segmentB < segmentA+segmentC && segmentC < segmentA+segmentB {
		fmt.Print("It is possible to form a triangle with the given segments.")
	} else {
		fmt.Print("It is not possible to form a triangle with the given segments.")
	}

}
