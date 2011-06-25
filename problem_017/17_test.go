package main

import (
	"testing"
)

type Problem17Test struct {
	in int
	out string
}

var Problem17Tests = []Problem17Test{
	Problem17Test{1, "one"},
	Problem17Test{2, "two"},
	Problem17Test{3, "three"},
	Problem17Test{4, "four"},
	Problem17Test{5, "five"},
	Problem17Test{6, "six"},
	Problem17Test{7, "seven"},
	Problem17Test{8, "eight"},
	Problem17Test{9, "nine"},
}

func TestProblem17(t *testing.T) {
	for _, v := range Problem17Tests {
		expected := v.out
		got := WriteNumber(v.in)
		if expected != got {
			t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
		}
	}
}

