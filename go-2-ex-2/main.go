package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	// Berechnung der Fibonacci-Zahlen bis zum Index 4
	fibs[2] = fibs[0] + fibs[1] // 2
	fibs[3] = fibs[1] + fibs[2] // 3
	fibs[4] = fibs[2] + fibs[3] // 5

	// Nächste Fibonacci-Zahl berechnen und anhängen
	fibs = append(fibs, fibs[3] + fibs[4]) // 8
	// Berechnung der nächsten drei Fibonacci-Zahlen
	fibs = append(fibs, fibs[4] + fibs[5]) // 13
	fibs = append(fibs, fibs[5] + fibs[6]) // 21
	fibs = append(fibs, fibs[6] + fibs[7]) // 34

	fmt.Println(fibs) // erwartete Ausgabe: [1 1 2 3 5 8 13 21 34]
}