package main

import (
	"testing"
)

type SumPrimesTest struct {
	in float64
	out uint64
}

var	SumPrimesTests = []SumPrimesTest{
	SumPrimesTest{10, 17},
	SumPrimesTest{100, 1060},
	SumPrimesTest{200, 4227},
	SumPrimesTest{543, 24133},
}

func TestSumPrimes(t *testing.T) {
	for _, v := range SumPrimesTests {
		got := SumPrimes(v.in)
		if v.out != got {
			t.Errorf("Sum of primes below %q should be %q, but was %q.", v.in, v.out, got)
		}
	}
}
