package main

import (
	"testing"
)

type ExtractDecimalTest struct {
	in float64
	out string
}

var ExtractDecimalTests = []ExtractDecimalTest{
	ExtractDecimalTest{ 0.1240, "124" },
	ExtractDecimalTest{ 5.24565, "24565" },
	ExtractDecimalTest{ 5.245655627147, "245655627147" },
}

func TestExtractDecimal(t *testing.T) {
	for _, test := range ExtractDecimalTests {
		expected := test.out
		got := ExtractDecimal(test.in)
		if expected != got {
			t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
		}
	}
}


