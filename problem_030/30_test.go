package main

import (
	"testing"
)

func TestIntPow(t *testing.T) {
	expected := 8
	got := IntPow(2, 3)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}


func TestIsSumOfExponents(t *testing.T) {
	if !IsSumOfExponents(1634, 4) {
		t.Errorf("1634 is a sum of fourth powers!")
	}
}

func TestSumOfPowersOfDigits(t *testing.T) {
	expected := 19316
	got := SumOfPowersOfDigits(4)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}
