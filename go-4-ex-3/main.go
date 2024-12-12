package main

import (
	"fmt"
	"math"
)

// computeQuadraticFormula berechnet die Lösungen einer quadratischen Gleichung.
func computeQuadraticFormula(a, b, c float64) []float64 {
	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant > 0 {
		// Zwei Lösungen
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		return []float64{x1, x2}
	} else if discriminant == 0 {
		// Eine Lösung
		x := -b / (2 * a)
		return []float64{x}
	} else {
		// Keine Lösung
		return []float64{}
	}
}

func main() {
	// Test 1: Zwei Lösungen (D > 0)
	a1, b1, c1 := 3.0, 4.0, 1.0
	fmt.Printf("Test 1 (D > 0): Lösungen = %v\n", computeQuadraticFormula(a1, b1, c1))

	// Test 2: Eine Lösung (D = 0)
	a2, b2, c2 := 2.0, 4.0, 2.0
	fmt.Printf("Test 2 (D = 0): Lösungen = %v\n", computeQuadraticFormula(a2, b2, c2))

	// Test 3: Keine Lösung (D < 0)
	a3, b3, c3 := 3.0, 4.0, 2.0
	fmt.Printf("Test 3 (D < 0): Lösungen = %v\n", computeQuadraticFormula(a3, b3, c3))
}
