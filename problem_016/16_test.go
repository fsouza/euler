package main

import (
	"testing"
)

func TestSumDigitsOfANumber(t *testing.T) {
	expected := uint64(9)
	got := SumDigits(27)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %d.\nGot: %d.", expected, got)
	}
}

func TestPow(t *testing.T) {
	expected := uint64(4)
	got := Pow(2, 2)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %d.\nGot: %d.", expected, got)
	}
}

