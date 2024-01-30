/*9) Faça um algoritmo que leia quanto dinheiro uma pessoa tem na carteira (em R$)
e mostre quantos dólares ela pode comprar. Considere US$1,00 = R$3,45.*/

package main

import "fmt"

func main() {
	var moneyInWallet float64
	var amountInDollar float64

	fmt.Print("Enter an amount in real currency: ")
	fmt.Scan(&moneyInWallet)

	amountInDollar = moneyInWallet / 3.45

	fmt.Printf("You can buy %.2f dollars", amountInDollar)

}
