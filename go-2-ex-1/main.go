package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	Name            FullName
	Birth           BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		Name: FullName{
			FirstName: "Leandro",
			LastName:  "Schwegler",
		},
		Birth: BirthDate{
			Day:   12,
			Month: 1,
			Year:  2008,
		},
		NumberOfSiblings: 1,
		ZodiacSign:  '\u2651', 
	}

	fmt.Println("Profile:", me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)

	me.NumberOfSiblings++

	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
