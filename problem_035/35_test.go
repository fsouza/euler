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
	primes := BuildPrimesList(1000)

	for _, test := range IsCircularTests {
		expected := test.out
		got := IsCircular(test.in, primes)
		if expected != got {
			t.Errorf("Assertion error.\nExpected %v to be circular? %v.\nGot: %v.", test.in, expected, got)
		}
	}
}

func TestBuildPrimesList(t *testing.T) {
	expected := []int{2, 3, 5, 7}
	got := BuildPrimesList(10)
	for i, prime := range got {
		if prime != expected[i] {
			t.Errorf("The primes list was built wrong")
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

func TestIsPresent(t *testing.T) {
	expected := true
	got := IsPresent([]int{ 1, 2, 3}, 3)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

