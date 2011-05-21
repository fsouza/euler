package main

import (
	"testing"
)

type FactorTest struct {
	in uint64
	out uint64
}

var FactorTests = []FactorTest{
	FactorTest{3, 3},
	FactorTest{9, 3},
	FactorTest{18, 3},
	FactorTest{35, 7},
	FactorTest{500, 5},
	FactorTest{2238, 373},
}

func TestFactor(t *testing.T) {
	for _, v := range FactorTests {
		got := Factor(v.in)
		if got != v.out {
			t.Errorf("%q's factor must be %q, but was %q", v.in, v.out, got)
		}
	}
}
