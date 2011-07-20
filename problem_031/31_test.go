package main

import (
	"testing"
)

type WaysTest struct {
	in, out int
}

var WaysTests = []WaysTest{
	WaysTest{ 2, 2, },
	WaysTest{ 5, 4, },
}

func TestWays(t *testing.T) {
	for _, test := range WaysTests {
		expected := test.out
		got := CalculateDifferentWays(test.in)
		if expected != got {
			t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
		}
	}
}

func TestDivMod(t *testing.T) {
	expectedDiv, expectedMod := 5, 2
	gotDiv, gotMod := DivMod(17, 3)

	if expectedDiv != gotDiv {
		t.Errorf("Assertion error, expected div to be %v, but was %v.", expectedDiv, gotDiv)
	}

	if expectedMod != gotMod {
		t.Errorf("Assertion error, expected mod to be %v, but was %v.", expectedMod, gotMod)
	}
}

