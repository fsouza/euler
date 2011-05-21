package main

import (
	"testing"
)

type PowDifferenceTest struct {
	in, out int
}

var PowDifferenceTests = []PowDifferenceTest{
	PowDifferenceTest{2, 4},
	PowDifferenceTest{3, 22},
	PowDifferenceTest{4, 70},
	PowDifferenceTest{5, 170},
	PowDifferenceTest{6, 350},
	PowDifferenceTest{10, 2640},
}

func TestPowDifference(t *testing.T) {
	for _, v := range PowDifferenceTests {
		got := PowDifference(v.in)
		if got != v.out {
			t.Errorf("The PowDifference for %q should be %q, but was %q!", v.in, v.out, got)
		}
	}
}

