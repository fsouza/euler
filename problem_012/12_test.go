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
	expectedPrimes := new(vector.IntVector)
	for _, integer := range []int{2, 3, 5} {
		expectedPrimes.Push(integer)
	}

	expectedCounts := new(vector.IntVector)
	for _, integer := range []int{1, 2, 1} {
		expectedCounts.Push(integer)
	}

	gotPrimes, gotCounts := GetPrimeFactors(90)
	if !AreTwoIntVectorEquals(expectedPrimes, gotPrimes) || !AreTwoIntVectorEquals(expectedCounts, gotCounts) {
		t.Errorf("Expected primes %q. Got primes %q\nExpected count %q. Got count %q", expectedPrimes, gotPrimes)
	}
}
