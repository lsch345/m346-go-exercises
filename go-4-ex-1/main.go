package main

import (
	"fmt"
)
func computeGrade(gotPoints, maxPoints float64) float64 {
	return 1.0 + (5.0 * gotPoints / maxPoints)
}

func main() {
	fmt.Println(computeGrade(17.5, 28.0))
	fmt.Println(computeGrade(28.0, 28.0))
	fmt.Println(computeGrade(0.0, 28.0))
	fmt.Println(computeGrade(14.0, 28.0))
}
