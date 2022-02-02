package fibonacci_test

import (
	"github.com/stretchr/testify/assert"
	"go10/fibonacci"
	"testing"
)

var result uint64

func BenchmarkFibRecurse(b *testing.B) {
	var res uint64
	for i := 0; i < b.N; i++ {
		fib := fibonacci.FibRecurse()
		res = fib(20, false)
	}
	result = res
}

func BenchmarkFibArray(b *testing.B) {
	var res uint64
	for i := 0; i < b.N; i++ {
		res = fibonacci.FibArray(20)
	}
	result = res
}

func TestFibTestify(t *testing.T) {
	fib := fibonacci.FibRecurse()
	assert.Equal(t, uint64(144), fib(12, false))
	assert.Equal(t, uint64(144), fibonacci.FibArray(12))
}
