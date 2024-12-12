package main

import (
	"fmt"
)

func convertCelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func convertFahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	// Celsius to Fahrenheit
	c1 := 0.0
	c2 := 25.0
	c3 := -10.0

	f1 := convertCelsiusToFahrenheit(c1)
	f2 := convertCelsiusToFahrenheit(c2)
	f3 := convertCelsiusToFahrenheit(c3)

	fmt.Printf("%v°C = %v°F\n", c1, f1) // Erwartet: 0°C = 32°F
	fmt.Printf("%v°C = %v°F\n", c2, f2) // Erwartet: 25°C = 77°F
	fmt.Printf("%v°C = %v°F\n", c3, f3) // Erwartet: -10°C = 14°F

	// Fahrenheit to Celsius
	fmt.Printf("%v°F = %v°C\n", f1, convertFahrenheitToCelsius(f1)) // Erwartet: 32°F = 0°C
	fmt.Printf("%v°F = %v°C\n", f2, convertFahrenheitToCelsius(f2)) // Erwartet: 77°F = 25°C
	fmt.Printf("%v°F = %v°C\n", f3, convertFahrenheitToCelsius(f3)) // Erwartet: 14°F = -10°C
