package main

import (
	"testing"
)

func TestFindLongestRecurring(t *testing.T) {
	expected := 7
	got := FindLongestRecurring(11)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

