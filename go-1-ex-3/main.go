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

	//Ausgabe Würfelzahl und Zeit
	fmt.Fprintln(os.Stdout, "The dice shows:", eyes)
	fmt.Fprintln(os.Stderr, "The dice was rolled at:", when.Format("2006-01-02 15:04:05"))

	//Daten ins File schreiben
	eyesFile, _ := os.Create("eyes.txt")
	defer eyesFile.Close()
	fmt.Fprintf(eyesFile, "The dice shows: %d eyes\n", eyes)

	//Daten ins File schreiben
	logFile, _ := os.Create("dice.log")
	defer logFile.Close()
	fmt.Fprintf(logFile, "The dice was rolled at: %s\n", when.Format("2006-01-02 15:04:05"))
}