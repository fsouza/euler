package main

import (
	"testing"
)

func TestSumDigitsOfANumber(t *testing.T) {
	expected := 6
	got := SumDigits(123)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

func TestSumOfPowersOfDigits(t *testing.T) {
	expected := 19316
	got := SumOfPowersOfDigits(4)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

