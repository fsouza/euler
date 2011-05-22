package main

import (
	"testing"
)

type PrimeAtTest struct {
	in, out int
}

var PrimeAtTests = []PrimeAtTest{
	PrimeAtTest{2, 3},
	PrimeAtTest{3, 5},
	PrimeAtTest{5, 11},
	PrimeAtTest{10, 29},
	PrimeAtTest{100, 541},
}

func TestPrimeAt(t *testing.T) {
	for _, v := range PrimeAtTests {
		got := PrimeAt(v.in)
		if v.out != got {
			t.Errorf("The %dth prime must be %d, but was %d", v.in, v.out, got)
		}
	}
}
