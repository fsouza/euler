package main

import (
	"testing"
)

func AssertEqual(expected, got int, t *testing.T) {
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

func TestCanFindTheGreatestSum(t *testing.T) {
	rows := []Row{
		Row{1},
		Row{2, 3},
		Row{4, 5, 6},
	}

	expected := 10
	got := FindGreatestSum(rows)
	AssertEqual(expected, got, t)
}

