package main

import (
	"fmt"
	"os"
)

func main() {
	var num int

	fmt.Print("Enter number (positive and greater then 1): ")
	fmt.Scanln(&num)

	if num <= 0 {
		fmt.Println("Number must be positive")
		os.Exit(1)
	} else if num == 1 {
		fmt.Println("There is no prime numbers between 0 and 1")
		os.Exit(1)
	} else if num == 2 {
		fmt.Println("Prime nunmers between 0 and 2: \n2")
		os.Exit(1)
	} else {
		fmt.Printf("Prime nunmers between 0 and %d \n2\n", num)
	}

	for i := 3; i <= num; i += 2 {
		isPrime := true
		for j := 3; j < i; j += 2 {
			if i%j == 0 {
				isPrime = false
			}
		}
		if isPrime == true {
			fmt.Println(i)
		}
	}
}
