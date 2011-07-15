package main

import (
	"testing"
)

type IsAbundantTest struct {
	in uint
	out bool
}

var IsAbundantTests = []IsAbundantTest{
	IsAbundantTest{ 12, true },
	IsAbundantTest{ 15, false },
	IsAbundantTest{ 24, true },
}

func TestCanCheckIfANumberIsAbundant(t *testing.T) {
	for _, test := range IsAbundantTests {
		expected := test.out
		got := IsAbundant(test.in)
		if expected != got {
			t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
		}
	}
}

