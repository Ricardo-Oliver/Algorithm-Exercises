/*11) Desenvolva uma lógica que leia os valores de A, B e C de uma equação do
segundo grau e mostre o valor de Delta.*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var c float64
	var delta float64

	fmt.Print("Enter the value A: ")
	fmt.Scan(&a)

	fmt.Print("Enter the value B: ")
	fmt.Scan(&b)

	fmt.Print("Enter the value C: ")
	fmt.Scan(&c)

	delta = math.Pow(b, 2) - (4 * a * c)

	fmt.Printf("The value delta is %v", delta)
}
