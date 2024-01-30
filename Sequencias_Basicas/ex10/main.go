/*10) Faça um algoritmo que leia a largura e altura de uma parede, calcule e
mostre a área a ser pintada e a quantidade de tinta necessária para o serviço,
sabendo que cada litro de tinta pinta uma área de 2metros quadrados.*/

package main

import "fmt"

func main() {
	var height int
	var width int
	var area int
	var paintedArea int

	fmt.Print("Enter the wall width: ")
	fmt.Scan(&width)

	fmt.Print("Enter the wall height: ")
	fmt.Scan(&height)

	area = height * width
	paintedArea = area / 2

	fmt.Printf("The area to be painted is approximately %v square meters and will require %v liters of paint.", area, paintedArea)
}
