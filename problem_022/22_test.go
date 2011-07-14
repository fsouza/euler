package main

import (
	"testing"
)

func TestCanFindTheSumOfAWord(t *testing.T) {
	expected := 53
	got := CountLetterNumbers("COLIN")
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}
