package main

import "fmt"

func main() {
	var s int16

	fmt.Print("Enter 3-digit number: ")
	fmt.Scanln(&s)

	if s < 0 {
		s = s * (-1)
	}

	var h = s / 100
	var u = s % 10
	var d = (s%100 - u) / 10

	fmt.Printf("Hundreds: %d\n", h)
	fmt.Printf("Dozens: %d\n", d)
	fmt.Printf("Units: %d\n", u)
}
