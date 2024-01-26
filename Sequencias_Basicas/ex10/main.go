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
