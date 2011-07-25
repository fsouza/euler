package main

import (
	"testing"
)

type IsPandigitalTest struct {
	in string
	out bool
}

var IsPandigitalTests = []IsPandigitalTest{
	IsPandigitalTest{ "123", false, },
	IsPandigitalTest{ "123456789", true, },
	IsPandigitalTest{ "192384576", true, },
	IsPandigitalTest{ "192984576", false, },
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

