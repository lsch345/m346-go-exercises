package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	// Testfälle
	a1, b1 := 3.0, 4.0
	fmt.Printf("Katheten: a=%.2f, b=%.2f, Hypotenuse: c=%.2f\n", a1, b1, computeHypotenuse(a1, b1))

	a2, b2 := 5.0, 12.0
	fmt.Printf("Katheten: a=%.2f, b=%.2f, Hypotenuse: c=%.2f\n", a2, b2, computeHypotenuse(a2, b2))

	a3, b3 := 8.0, 15.0
	fmt.Printf("Katheten: a=%.2f, b=%.2f, Hypotenuse: c=%.2f\n", a3, b3, computeHypotenuse(a3, b3))

	a4, b4 := 1.0, 1.0
	fmt.Printf("Katheten: a=%.2f, b=%.2f, Hypotenuse: c=%.2f\n", a4, b4, computeHypotenuse(a4, b4))
}