package main

import (
	"testing"
)

func AssertEqual(expected, got int, t *testing.T) {
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

func AssertAreSlicesEqual(expected, got []Row, t *testing.T) {
	if len(expected) != len(got) {
		t.Errorf("The two slices don't have the same size")
	}

	for i := 0; i < len(expected); i++ {
		expectedRow := expected[i]
		gotRow := got[i]
		if len(expectedRow) != len(gotRow) {
			t.Errorf("The rows don't have the same size!")
		}

		for j := 0; j < len(expectedRow); j++ {
			if expectedRow[j] != gotRow[j] {
				t.Errorf("The two slices are not equal!\nExpected: %v\nGot: %v", expected, got)
			}
		}
	}
}

func TestCanFindGreatestValueOfASlice(t *testing.T) {
	in := []int{1, 2, 10}
	expected := 10
	got := Max(in)
	AssertEqual(expected, got, t)
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

func TestBuildRowsSliceFromString(t *testing.T) {
	input := "1\n2 3\n4 5 6"

	expected := []Row{
		Row{1},
		Row{2, 3},
		Row{4, 5, 6},
	}

	got := BuildrowsSliceFromString(input)
	AssertAreSlicesEqual(expected, got, t)
}

