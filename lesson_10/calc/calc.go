package calc

import "fmt"

func Calc(op string, a, b float64) (result float64, err error) {
	result = 0
	err = nil

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			err = fmt.Errorf("cant divide by zero")
		}
		result = a / b
	default:
		err = fmt.Errorf("unknown operation")
	}
	return result, err
}
