package main

import "fmt"

var cache = map[int]int{0: 0, 1: 1}

func fibonacci(n int) int {
	if val, ok := cache[n]; ok {
		return val
	}
	cache[n] = fibonacci(n-1) + fibonacci(n-2)
	return cache[n]
}

func main() {
	var num int
	fmt.Print("Enter number: ")
	fmt.Scanln(&num)

	res := fibonacci(num)
	fmt.Printf("F(%d) = %d", num, res)
}
