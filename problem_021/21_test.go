package main

import (
	"testing"
)

type FactorsTest struct {
	in, out uint
}

var	FactorsTests = []FactorsTest{
	FactorsTest{ 220, 284 },
	FactorsTest{ 284, 220 },
}

func TestFindFactors(t *testing.T) {
	for _, test := range FactorsTests {
		expected := test.out
		got := SumOfFactors(test.in)
		if expected != got {
			t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
		}
	}
}

