package main

import (
	"testing"
)

func TestSumDigitsOfANumber(t *testing.T) {
	expected := uint(9)
	got := SumDigitsOf(27)
	if expected != got {
		t.Errorf("%d != %d", expected, got)
	}
}

