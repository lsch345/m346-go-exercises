package main

import "fmt"

func main() {
	// Eine Map erstellen, die Modulnummern mit ihren Beschreibungen verknüpft
	modules := map[int]string{
		104: "Programmieren",
		117: "Datenbanken",
		346: "Webentwicklung",
	}

	// Modul-Beschreibungen ausgeben
	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// Ein Modul löschen
	delete(modules, 117)

	// Ein neues Modul hinzufügen
	modules[200] = "KI"

	// Ein Modul ersetzen
	modules[104] = "Cloud"

	// Aktuellen Zustand der Map ausgeben
	fmt.Println(modules)
}