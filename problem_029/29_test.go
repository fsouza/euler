package main

import (
	"math/big"
	"testing"
)

func TestIsPresent(t *testing.T) {
	sequence := make([]*big.Int, 0)
	sequence = AddToSequence(sequence, big.NewInt(0), big.NewInt(1))

	if !IsPresent(sequence, big.NewInt(0)) {
		t.Errorf("Should find 0 in the sequence")
	}
}

func TestAddToSequence(t *testing.T) {
	sequence := make([]*big.Int, 0)
	sequence = AddToSequence(sequence, big.NewInt(0), big.NewInt(1))
	sequence = AddToSequence(sequence, big.NewInt(0))

	if len(sequence) != 2 {
		t.Errorf("Length of sequence should be 2, but was %d.", len(sequence))
	}
}

func TestFindTermsCountInSequence(t *testing.T) {
	expected := 15
	got := CountTermsInSequence(5, 5)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}
