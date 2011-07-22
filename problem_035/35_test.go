package main

import (
	"testing"
)

type IsCircularTest struct {
	in int
	out bool
}

var IsCircularTests = []IsCircularTest {
	IsCircularTest{ 7, true, },
	IsCircularTest{ 17, true, },
	IsCircularTest{ 19, false, },
	IsCircularTest{ 23, false, },
	IsCircularTest{ 40, false, },
	IsCircularTest{ 197, true, },
}

func TestIsCircular(t *testing.T) {
	for _, test := range IsCircularTests {
		expected := test.out
		got := IsCircular(test.in)
		if expected != got {
			t.Errorf("Assertion error.\nExpected %v to be circular? %v.\nGot: %v.", test.in, expected, got)
		}
	}
}

func TestRotate(t *testing.T) {
	expected := 971
	got := Rotate(197)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}
