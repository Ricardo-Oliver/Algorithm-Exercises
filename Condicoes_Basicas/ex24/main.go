/*24) Faça um algoritmo que pergunte a distância que um passageiro deseja
percorrer em Km. Calcule o preço da passagem, cobrando R$0.50 por Km para
viagens até 200Km e R$0.45 para viagens mais longas.*/

package main

import "fmt"

func main() {
	var distance int

	fmt.Print("Enter the distance you want to travel: ")
	fmt.Scan(&distance)

	if distance <= 200 {
		valueTravel := float64(distance) * 0.5
		fmt.Printf("The cost of the trip is U$%.2f", valueTravel)
		return
	}

	if distance > 200 {
		valueTravel := float64(distance) * 0.45
		fmt.Printf("The cost of the trip is U$%.2f", valueTravel)
	}
}
