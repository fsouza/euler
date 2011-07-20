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

func TestLenOfNumber(t *testing.T) {
	expected := 5
	got := LenOfNumber(12345)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

func TestBuildLimitsForLength(t *testing.T) {
	min, max := BuildLimitsForLength(4)
	if min != 1000 {
		t.Errorf("Expected mininum to be 1000, but was %v.", min)
	}

	if max != 9999 {
		t.Errorf("Expected maximun to be 1000, but was %v.", max)
	}
}


