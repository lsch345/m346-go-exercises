package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	eyes := rand.Intn(6) + 1 
	when := time.Now()

	file, _ := os.Create("output.txt")
	defer file.Close()

	fmt.Fprintln(file, "the dice shows", eyes, "eyes")
	fmt.Fprintln(file, "the dice was rolled at", when)

	fmt.Println("Ergebnisse wurden in output.txt gespeichert.")
}

