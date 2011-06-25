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

