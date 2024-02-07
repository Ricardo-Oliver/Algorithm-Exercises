/*17) Escreva um programa que pergunte a velocidade de um carro. Caso ultrapasse
80Km/h, exiba uma mensagem dizendo que o usuário foi multado. Nesse caso, exiba
o valor da multa, cobrando R$5 por cada Km acima da velocidade permitida.*/

package main

import "fmt"

func main() {
	var speedCar int
	var fineValue float64

	fmt.Println("What is the speed of the car ?")
	fmt.Scan(&speedCar)

	if speedCar > 80 {
		fineValue = (float64(speedCar) - 80) * 5
		fmt.Printf("You have been fined. In the amount of U$%.2f", fineValue)
		return
	}

	fmt.Print("Have a good trip")

}
