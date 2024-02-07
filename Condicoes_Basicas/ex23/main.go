/*23) Numa promoção exclusiva para o Dia da Mulher, uma loja quer dar descontos
para todos, mas especialmente para mulheres. Faça um programa que leia nome,
sexo e o valor das compras do cliente e calcule o preço com desconto. Sabendo
que:
- Homens ganham 5% de desconto
- Mulheres ganham 13% de desconto*/

package main

import "fmt"

func main() {
	var name string
	var gender string
	var buyValue float64

	fmt.Print("Enter the name: ")
	fmt.Scan(&name)

	fmt.Print("Enter the gender: ")
	fmt.Scan(&gender)

	fmt.Print("Enter the buy value: ")
	fmt.Scan(&buyValue)

	if gender == "Masculino" {
		totalValue := buyValue - (buyValue * 0.05)
		fmt.Printf("The total purchase price with 5%% discount is U$%.2f", totalValue)
	} else {
		totalValue := buyValue - (buyValue * 0.13)
		fmt.Printf("The total purchase price with 13%% discount is U$%.2f", totalValue)
	}
}
