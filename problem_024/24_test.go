package main

import (
	"testing"
)

func TestGenerateAllPermutations(t *testing.T) {
	expected := "102"
	got := FindNthPermutation(2, 0, 1, 2)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

func TestFactorial(t *testing.T) {
	expected := 120
	got := Factorial(5)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

