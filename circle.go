package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64

	fmt.Print("Enter circle area: ")
	fmt.Scanln(&s)

	if s <= 0 {
		fmt.Print("Area cant be negative or zero")
		return
	}

	var r = math.Sqrt(s / math.Pi)

	fmt.Printf("Circle diameter: %f\n", r*2)
	fmt.Printf("Circle length: %f\n", r*2*math.Pi)
}
