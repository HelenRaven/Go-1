package sorting_test

import (
	"go10/sorting"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestSortingTable(t *testing.T) {
	type test struct {
		have []int
		want []int
	}

	tests := []test{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
	}

	for _, c := range tests {
		result := sorting.Sorting(c.have)
		if !reflect.DeepEqual(result, c.want) {
			t.Fatalf("Expected: %v, recieved: %v", c.want, result)
		}
	}
}
func TestSortingRandom(t *testing.T) {
	want := []int{}

	for i := 0; i < 100; i++ {
		for j := rand.Intn(100); j > 0; j-- {
			want = append(want, rand.Int())
		}

		have := make([]int, len(want))
		copy(have, want)

		sort.Ints(want)
		have = sorting.Sorting(have)

		if !reflect.DeepEqual(have, want) {
			t.Fatalf("Expected: %v, recieved: %v", want, have)
		}
	}
}
