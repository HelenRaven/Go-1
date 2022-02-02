package calc_test

import (
	"fmt"
	"go10/calc"
	"testing"
)

func TestCalc(t *testing.T) {
	_, err := calc.Calc("/", 10, 0)
	if err == nil {
		t.Fatalf("Calc(+ 10 0) expected error, but not received it")
	}

	_, err = calc.Calc("max", 10, 20)
	if err == nil {
		t.Fatalf("Calc(max 10 20) expected error, but not received it")
	}
}

func TestCalcTable(t *testing.T) {
	type test struct {
		op         string
		a, b, want float64
	}

	tests := []test{
		{"+", 10, 0, 10},
		{"+", 0, 3, 3},
		{"+", -5, 0, -5},
		{"+", 10, 45, 55},

		{"-", 10, 0, 10},
		{"-", 0, 3, -3},
		{"-", -5, 0, -5},
		{"-", 10, 45, -35},

		{"*", 10, 0, 0},
		{"*", 0, 3, 0},
		{"*", -5, 1, -5},
		{"*", 10, 45, 450},

		{"/", 10, 1, 10},
		{"/", 9, 3, 3},
		{"/", -5, 1, -5},
		{"/", 10, -5, -2},
	}
	for _, c := range tests {
		culcStr := fmt.Sprintf("Culc( %f %s %f", c.a, c.op, c.b)

		result, err := calc.Calc(c.op, c.a, c.b)
		if err != nil {
			t.Fatalf("%s expected no errors, but returned %s", culcStr, err)
		}

		if result != c.want {
			t.Fatalf("Expected: %f, recieved: %f", c.want, result)
		}
	}

}
