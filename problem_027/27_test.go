package main

import (
	"testing"
)

func AssertAreSliceEqual(expected, got []int, t *testing.T) {
	if len(expected) != len(got) {
		t.Errorf("Assertion error.\nExpected: %v\nGot:%v", expected, got)
	} else {
		for i := 0; i < len(expected); i++ {
			if expected[i] != got[i] {
				t.Errorf("Assertion error.\nExpected: %v\nGot:%v", expected, got)
			}
		}
	}
}

func TestBuildSliceOfPrimes(t *testing.T) {
	expected := []int{ 2, 3, 5, 7, }
	got := BuildSliceOfPrimes(10)
	AssertAreSliceEqual(expected, got, t)
}

type IsPrimeTest struct {
	in int
	out bool
}

var IsPrimeTests = []IsPrimeTest{
	IsPrimeTest{ 2, true },
	IsPrimeTest{ 3, true },
	IsPrimeTest{ 5, true },
	IsPrimeTest{ 10, false },
	IsPrimeTest{ 505, false },
	IsPrimeTest{ -505, false},
	IsPrimeTest{ -2, true},
}

func TestIsPrime(t *testing.T) {
	for _, test := range IsPrimeTests {
		expected := test.out
		got := IsPrime(test.in)
		if expected != got {
			t.Errorf("Assertion error.\nExpected %v to be prime? %v.\nGot: %v.", test.in, expected, got)
		}
	}
}

