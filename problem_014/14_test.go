package main

import (
	"testing"
)

func TestFindChainLengthForNumber(t *testing.T) {
	expected := 11
	got := FindChainLength(26)
	if expected != got {
		t.Errorf("%d != %d", expected, got)
	}
}

