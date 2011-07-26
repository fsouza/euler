package main

import (
	"testing"
)

func TestFindNthDigitOfTheFractionalPart(t *testing.T) {
	expected := uint8(1)
	got := FindNthDigit(12)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

