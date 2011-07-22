package main

import (
	"testing"
)

type ReverseTest struct {
	in, out []int
}

var ReverseTests = []ReverseTest{
	ReverseTest{ []int{ 48, 49, 50, }, []int{ 50, 49, 48, } },
	ReverseTest{ []int{ 48, 49, 0, }, []int{ 49, 48, } },
}

func TestReverse(t *testing.T) {
	for _, test := range ReverseTests {
		expected := test.out
		got := Reverse(test.in)
		if string(expected) != string(got) {
			t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
		}
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
