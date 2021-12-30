package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputNums := []int64{}
	var temp int64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		inputNums = append(inputNums, num)
	}

	for i := 1; i < len(inputNums); i++ {
		temp = inputNums[i]
		for j := i - 1; j >= 0; j-- {
			if inputNums[j] < temp {
				continue
			} else {
				inputNums[j+1] = inputNums[j]
				inputNums[j] = temp
			}
		}
	}
	fmt.Println(inputNums)
}
