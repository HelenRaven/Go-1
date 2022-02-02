package sorting

func Sorting(inputNums []int) []int {
	var temp int

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
	return inputNums
}
