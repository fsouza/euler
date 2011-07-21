package main

import (
	"testing"
)

type IsNonTrivialTest struct {
	numerator, denominator int
	out bool
}

var	IsNonTrivialTests = []IsNonTrivialTest {
	IsNonTrivialTest{ 30, 50, false, },
	IsNonTrivialTest{ 49, 98, true, },
}

func TestIsNonTrivial(t *testing.T) {
	for _, test := range IsNonTrivialTests {
		expected := test.out
		got := IsNonTrivial(test.numerator, test.denominator)
		if expected != got {
			t.Errorf("Assertion error.\nExpected %v / %v to be trivial? %v.\nGot: %v.", test.numerator, test.denominator, expected, got)
		}
	}
}

