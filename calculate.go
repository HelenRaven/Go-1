package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	fmt.Print("Enter operation (+, -, *, / , % , min , max ): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Cannot be divided by zero!")
			os.Exit(1)
		} else {
			res = a / b
		}
	case "%":
		if b == 0 {
			fmt.Println("Cannot be divided by zero!")
			os.Exit(1)
		} else {
			res = math.Mod(a, b)
		}
	case "min":
		res = math.Min(a, b)
	case "max":
		res = math.Max(a, b)
	default:
		fmt.Println("Wrong operation!")
		os.Exit(1)
	}

	fmt.Printf("Result: %f\n", res)
}
