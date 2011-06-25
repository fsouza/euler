package main

import (
	"testing"
)

type Problem17Test struct {
	in, out int
}

var Problem17Tests = []Problem17Test{
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
	Problem17Test{20, 6},
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
			t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
		}
	}
}

