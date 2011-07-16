package main

import (
	"testing"
)

func AssertAreSlicesEqual(expected, got []string, t *testing.T) {
	if len(expected) != len(got) {
		t.Errorf("Expected: %v\nGot: %v", expected, got)
	} else {
		for i := 0; i < len(expected); i++ {
			if !t.Failed() && expected[i] != got[i] {
				t.Errorf("Expected: %v\nGot: %v", expected, got)
			}
		}
	}
}

func TestGenerateAllPermutations(t *testing.T) {
	expected := []string{ "012", "021", "102", "120", "201", "210", }
	got := GeneratePermutations(0, 1, 2)
	AssertAreSlicesEqual(expected, got, t)
}

func TestFactorial(t *testing.T) {
	expected := 120
	got := Factorial(5)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

