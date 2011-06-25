package main

import (
	"testing"
)

type Problem17Test struct {
	in, out int
}

var Problem17Tests = []Problem17Test{
	Problem17Test{0, 0},
	Problem17Test{1, 3},
	Problem17Test{2, 3},
	Problem17Test{3, 5},
	Problem17Test{4, 4},
	Problem17Test{5, 4},
	Problem17Test{6, 3},
	Problem17Test{7, 5},
	Problem17Test{8, 5},
	Problem17Test{9, 4},
	Problem17Test{10, 3},
	Problem17Test{11, 6},
	Problem17Test{12, 6},
	Problem17Test{13, 8},
	Problem17Test{14, 8},
	Problem17Test{15, 7},
	Problem17Test{20, 6},
	Problem17Test{21, 9},
	Problem17Test{22, 9},
	Problem17Test{23, 10},
	Problem17Test{29, 10},
	Problem17Test{30, 6},
	Problem17Test{31, 9},
	Problem17Test{40, 6},
	Problem17Test{50, 5},
	Problem17Test{60, 5},
	Problem17Test{70, 7},
	Problem17Test{80, 6},
	Problem17Test{90, 6},
}

func TestDivMod(t *testing.T) {
	expectedDiv, expectedMod := 2, 1
	gotDiv, gotMod := DivMod(5, 2)
	if expectedDiv != gotDiv || expectedMod != gotMod {
		t.Errorf("Assertion error.\n%v != %v or %v != %v", expectedDiv, gotDiv, expectedMod, gotMod)
	}
}


func TestProblem17(t *testing.T) {
	for _, v := range Problem17Tests {
		expected := v.out
		got := CountLetters(v.in)
		if expected != got {
			t.Errorf("Assertion error.\nExpected: %v for %v.\nGot: %v.", expected, v.in, got)
		}
	}
}

