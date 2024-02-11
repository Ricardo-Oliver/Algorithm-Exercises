/*30) [DESAFIO] Refaça o algoritmo 25, acrescentando o recurso de mostrar que tipo
de triângulo será formado:
- EQUILÁTERO: todos os lados iguais
- ISÓSCELES: dois lados iguais
- ESCALENO: todos os lados diferentes*/

package main

import "fmt"

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
		if segmentA == segmentB && segmentB == segmentC {
			fmt.Println("Type of triangle formed: Equilateral")
		} else if segmentA == segmentB || segmentA == segmentC || segmentB == segmentC {
			fmt.Println("Type of triangle formed: Isosceles")
		} else {
			fmt.Println("Type of triangle formed: Scalene")
		}
		fmt.Print("It is possible to form a triangle with the given segments.")
	} else {
		fmt.Print("It is not possible to form a triangle with the given segments.")
	}
}
