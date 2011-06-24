package main

import (
	"testing"
)

func TestFindChainLengthForNumber(t *testing.T) {
	expected := 10
	got := FindChainLength(13)
	if expected != got {
		t.Errorf("%d != %d", expected, got)
	}
}

