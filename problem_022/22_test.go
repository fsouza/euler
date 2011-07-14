package main

import (
	"testing"
)

func TestCanFindTheSumOfAWord(t *testing.T) {
	expected := int64(53)
	got := CountLetterNumbers("COLIN")
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

func TestCanFindSumOfStringOnlyCountingLetters(t *testing.T) {
	expected := int64(53)
	got := CountLetterNumbers("CO##LIN\n")
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

