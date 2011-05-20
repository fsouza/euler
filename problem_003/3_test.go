package main

import (
	"testing"
)

type PrimeTest struct {
	number uint
	isPrime bool
}

var PrimeTests = []PrimeTest{
	PrimeTest{2, true},
	PrimeTest{5, true},
	PrimeTest{7, true},
	PrimeTest{8, false},
	PrimeTest{13, true},
	PrimeTest{69, false},
}

func TestPrime(t* testing.T) {
	for _, v := range PrimeTests {
		got := IsPrime(v.number)
		if got != v.isPrime {
			t.Errorf("Is %q prime? The response shoud be %q, but was %q.", v.number, v.isPrime, got)
		}
	}
}

type FactorTest struct {
	in float64
	out uint
}

var FactorTests = []FactorTest{
	FactorTest{2, 2},
	FactorTest{3, 3},
	FactorTest{8, 2},
	FactorTest{9, 3},
	FactorTest{18, 3},
	FactorTest{35, 7},
	FactorTest{500, 5},
}

func TestFactor(t *testing.T) {
	for _, v := range FactorTests {
		got := Factor(v.in)
		if got != v.out {
			t.Errorf("%q != %q", v.out, got)
		}
	}
}
