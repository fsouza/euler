package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	expected := "123"
	out := Reverse([]int("321"))
	got := string(out)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

type ToBinaryTest struct {
	in int
	out string
}

var ToBinaryTests = []ToBinaryTest{
	ToBinaryTest{ 1, "1", },
	ToBinaryTest{ 2, "10", },
	ToBinaryTest{ 3, "11", },
	ToBinaryTest{ 5, "101", },
	ToBinaryTest{ 15, "1111", },
}

func TestToBinary(t *testing.T) {
	expected := "110"
	got := ToBinary(6)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

func TestDivMod(t *testing.T) {
	expectedDiv, expectedMod := 2, 1
	gotDiv, gotMod := DivMod(5, 2)
	if expectedDiv != gotDiv {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expectedDiv, gotDiv)
	}

	if expectedMod != gotMod {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expectedMod, gotMod)
	}
}

func TestIsTheSameReversed(t *testing.T) {
	expected := true
	got := IsTheSameReversed("asa")
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}
