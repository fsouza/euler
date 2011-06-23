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

func TestSliceDigitsOfANumber(t *testing.T) {
	expected := "1234"
	got := SliceDigits(123456789, 4)
	if expected != got {
		t.Errorf("%s != %s", expected, got)
	}
}

