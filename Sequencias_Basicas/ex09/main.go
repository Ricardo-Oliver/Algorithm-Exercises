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
