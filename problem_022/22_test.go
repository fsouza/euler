package main

import (
	"big"
	"testing"
)

func TestCanFindTheSumOfAWord(t *testing.T) {
	expected := big.NewInt(53)
	got := CountLetterNumbers("COLIN")
	if expected.Int64() != got.Int64() {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}
