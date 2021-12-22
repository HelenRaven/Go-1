package main

import "fmt"

func main() {
	var a, b float32

	fmt.Print("Enter first side length: ")
	fmt.Scanln(&a)

	fmt.Print("Enter second side length: ")
	fmt.Scanln(&b)

	if a <= 0 || b <= 0 {
		fmt.Print("Sides length cant be negative or zero")
		return
	}

	fmt.Printf("Rectangle area: %f\n", a*b)
}
