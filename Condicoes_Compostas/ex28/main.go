/*28) Faça um programa que leia a largura e o comprimento de um terreno
retangular, calculando e mostrando a sua área em m². O programa também
devemostrar a classificação desse terreno, de acordo com a lista abaixo:
- Abaixo de 100m² = TERRENO POPULAR
- Entre 100m² e 500m² = TERRENO MASTER
- Acima de 500m² = TERRENO VIP.*/

package main

import "fmt"

func main() {
	var (
		width  float64
		length float64
		area   float64
	)

	fmt.Print("Enter the width of the land: ")
	fmt.Scan(&width)

	fmt.Print("Enter the length of the land: ")
	fmt.Scan(&length)

	area = width * length

	if area < 100 {
		fmt.Print("Popular terrain.")
	} else if area >= 100 && area <= 500 {
		fmt.Print("Master land.")
	} else {
		fmt.Print("VIP land.")
	}
}
