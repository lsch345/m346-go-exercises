package main

import "fmt"

func main() {

	//Variabeln deklarieren
	firstName := "Leandro"
	lastName := "Schwegler"
	dayOfBirth := 12
	monthOfBirth := 1
	yearOfBirth := 2008
	numberOfSiblings := 1
	heightInMeters := 1.80
	zodiacSign := '\u2651'

	//Variabeln ausprinten
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
