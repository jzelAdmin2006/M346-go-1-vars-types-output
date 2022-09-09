package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	var firstName string = "Johannes"
	var lastName string = "Zeller"
	var dayOfBirth int = 1
	var monthOfBirth int = 4
	var yearOfBirth int = 2006
	var numberOfSiblings int = 2
	var heightInMeters float32 = 1.85
	var zodiacSign rune = '\u2648'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
