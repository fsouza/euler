package main

import (
	"testing"
)

func TestFindDayOfWeek(t *testing.T) {
	expected := 5 // 5 == thursday
	got := FindDayOfWeek(1989, 2, 16)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

