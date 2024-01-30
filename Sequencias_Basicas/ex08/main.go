/*8) Desenvolva um programa que leia uma distância em metros e mostre os valores
relativos em outras medidas.
Ex:
Digite uma distância em metros: 185.72
A distância de 85.7m corresponde a:
0.18572Km
1.8572Hm
18.572Dam
1857.2dm
18572.0cm
185720.0mm*/

package main

import "fmt"

func main() {
	var meters float64

	fmt.Print("Enter a distance in meters: ")
	fmt.Scanln(&meters)

	km := meters / float64(1000)
	hm := meters / float64(100)
	dam := meters / float64(10)
	dm := meters * float64(10)
	cm := meters * float64(100)
	mm := meters * float64(1000)

	fmt.Printf("The distance of %.2fm corresponds the: \n%.5fkm\t\t\t%.1fdm\n%.4fhm\t\t\t%.1fcm\n%.3fdam\t\t\t%.1fmm", meters, km, dm, hm, cm, dam, mm)
}
