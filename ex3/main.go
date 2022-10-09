package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stderr, "the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	// go run ex3/main.go > eyes.txt 2> dice.log

	// Same file
	// go run ex3/main.go > dice.log 2>&1

	// Attach
	// go run ex3/main.go >> dice.log

	// Write in file and console
	// go run ex3/main | tee dice.log
}
