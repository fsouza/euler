package main

import (
	"testing"
)

func TestGetCellOnPascalTriangle(t *testing.T) {
	expected := 6
	got := GetPascalsTriangleCell(4, 2)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %d.\nGot: %d.", expected, got)
	}
}

