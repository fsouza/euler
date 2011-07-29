package main

import (
	"testing"
)

type IsPandigitalTest struct {
	in int
	out bool
}

var	IsPandigitalTests = []IsPandigitalTest{
	IsPandigitalTest{2134, true},
	IsPandigitalTest{2345, false},
	IsPandigitalTest{2445, false},
	IsPandigitalTest{123456789, true},
}

func TestIsPandigital(t *testing.T) {
	for _, test := range IsPandigitalTests {
		expected := test.out
		got := IsPandigital(test.in)
		if expected != got {
			t.Errorf("Assertion error.\nExpected %v to be pandigital? %v.\nGot: %v.", test.in, expected, got)
		}
	}
}

