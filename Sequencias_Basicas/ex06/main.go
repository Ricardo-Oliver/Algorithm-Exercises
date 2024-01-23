package main

import "fmt"

func main() {
	var number int
	var predecessor int
	var successor int

	fmt.Print("Enter a value: ")
	fmt.Scan(&number)

	predecessor = number - 1
	successor = number + 1

	fmt.Printf("The predecessor of %v is %v\n", number, predecessor)
	fmt.Printf("The successor of %v is %v", number, successor)
}
