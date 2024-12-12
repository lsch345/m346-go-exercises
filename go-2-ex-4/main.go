package main

import "fmt"

func main() {

	// Schüler mit Vorname und Nachname
	type Student struct {
		FirstName string
		LastName  string
	}

	// Klasse mit Liste von Schülern
	type Class struct {
		Students []Student
	}

	// Map Module mit allen Schülern
	modules := map[string]Class{
		"M113": {
			Students: []Student{
				{FirstName: "Noah", LastName: "Krummenacher"},
				{FirstName: "Timo", LastName: "Fuchs"},
				{FirstName: "Nils", LastName: "Steiner"},
			},
		},
		"M273": {
			Students: []Student{
				{FirstName: "Leandro", LastName: "Schwegler"},
				{FirstName: "Nando", LastName: "Schuermann"},
				{FirstName: "Tim", LastName: "Huesler"},
			},
		},
	}

	// Ausgabe
	fmt.Println("Modules and their classes:")
	for module, class := range modules {
		fmt.Println("Module:", module)
		for _, student := range class.Students {
			fmt.Printf("\tStudent: %s %s\n", student.FirstName, student.LastName)
		}
	}
}
