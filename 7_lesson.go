package lesson7

//1.С какими интерфейсами мы уже сталкивались в предыдущих уроках?
//Функции ввода/вывода из пакета fmt: Print(f)(ln) Scan(ln)

//2. Посмотрите примеры кода в своём портфолио. Везде ли ошибки обрабатываются грамотно? Хотите ли вы переписать какие-либо функции?

import (
	"errors"
	"fmt"
)

func prime(num int) error {
	if num <= 0 {
		return errors.New("Number must be positive")
	} else if num == 1 {
		return errors.New("There is no prime numbers between 0 and 1")
	} else if num == 2 {
		fmt.Printf("Prime nunmers between 0 and 2: \n2")
		return nil
	} else {
		fmt.Printf("Prime nunmers between 0 and %d \n2\n", num)
	}

	for i := 3; i <= num; i += 2 {
		isPrime := true
		for j := 3; j < i; j += 2 {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime == true {
			fmt.Println(i)
		}
	}
	return nil
}

////////////////////////////////////////////////////////////

type Operator string

const (
	OpPlus     Operator = "+"
	OpMinus    Operator = "-"
	OpMultiply Operator = "*"
	OpDivide   Operator = "/"
	OpRemains  Operator = "%"
)

func Calculate(a, b float32, op Operator) (res float32, err error) {
	switch op {
	case OpPlus:
		res = a + b
	case OpMinus:
		res = a - b
	case OpMultiply:
		res = a * b
	case OpDivide:
		if b == 0 {
			err = errors.New("Cant divide by zero")
		} else {
			res = a / b
		}
	}
	return res, err
}

func main() {
	var a, b float32
	var op Operator

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	fmt.Print("Enter operation (+, -, *, / ): ")
	fmt.Scanln(&op)

	res, err := Calculate(a, b, op)
	if err == nil {
		fmt.Println("Result: ", res)
	} else {
		fmt.Println(err.Error())
	}
}
