package main

import (
	"container/vector"
	"testing"
)

func AreTwoIntVectorEquals(vector1, vector2 *vector.IntVector) bool {
	if vector1.Len() != vector2.Len() {
		return false
	}

	for i := 0; i < vector1.Len(); i++ {
		if vector1.At(i) != vector2.At(i) {
			return false
		}
	}

	return true
}

func TestGetPrimeFactorsOfANumber(t *testing.T) {
	expected := new(vector.IntVector)
	for _, integer := range []int{2, 3, 5} {
		expected.Push(integer)
	}

	got := GetPrimeFactors(90)
	if !AreTwoIntVectorEquals(expected, got) {
		t.Errorf("Expected %q. Got %q", expected, got)
	}
}
