package main

import "fmt"

func main() {
	var priceProduct float64
	var pricePromotion float64

	fmt.Print("Enter the value of the product: U$")
	fmt.Scan(&priceProduct)

	pricePromotion = priceProduct - (priceProduct * 0.05)

	fmt.Printf("The promotional value of the product with 5%% discount is U$%.2f", pricePromotion)
}
