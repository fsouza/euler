package main

import (
	"math/big"
	"testing"
)

type FactorialTest struct {
	in  int64
	out *big.Int
}

var FactorialTests = []FactorialTest{
	FactorialTest{0, big.NewInt(1)},
	FactorialTest{1, big.NewInt(1)},
	FactorialTest{3, big.NewInt(6)},
	FactorialTest{4, big.NewInt(24)},
	FactorialTest{5, big.NewInt(120)},
}

func TestFactorial(t *testing.T) {
	for _, test := range FactorialTests {
		expected := test.out
		got := Factorial(test.in)
		if expected.Int64() != got.Int64() {
			t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
		}
	}
}
