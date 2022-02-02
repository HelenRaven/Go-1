package prime_test

import (
	"go10/prime"
	"reflect"
	"testing"
)

func TestPrime(t *testing.T) {
	_, err := prime.Prime(-10)
	if err == nil {
		t.Fatalf("Prime(-10) expected error, but not received it")
	}

	_, err = prime.Prime(1)
	if err == nil {
		t.Fatalf("Prime(1) expected error, but not received it")
	}

	res, err := prime.Prime(30)
	if err != nil {
		t.Fatalf("Prime(10) not expected error, but received: %s", err)
	}

	expect := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	if !reflect.DeepEqual(expect, res) {
		t.Fatalf("Expect: %v, recieved: %v", expect, res)
	}
}
