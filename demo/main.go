package main

import "fmt"

func main() {
	var result = 5
	var isReady = true
	var caption = "Hello"

	var a, b, c = 1, "Hey", false
	d, e, f := 2, "Ho", true

	const g float32 = 9.81
	const pi float64 = 3.14159265359
	const name = "Brandon"

	fmt.Println(result)
	fmt.Println(isReady)
	fmt.Println(caption)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
	fmt.Println(g, pi)
	fmt.Println(name)
}
