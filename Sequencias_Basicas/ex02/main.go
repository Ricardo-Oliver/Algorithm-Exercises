/*2) Faça um programa que leia o nome de uma pessoa e mostre uma mensagem de boasvindas
para ela:
Ex:
Qual é o seu nome? João da Silva
Olá João da Silva, é um prazer te conhecer!*/

package main

import "fmt"

func main() {
	var name string 

	fmt.Println("What's your name ?")
	fmt.Scan(&name)
	fmt.Printf("Hello %v, nice too meet you!", name)
}