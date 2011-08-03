package main

import (
	"testing"
)

func TestTriangleNumberAt(t *testing.T) {
	expected := 45
	got := TriangleNumberAt(9)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}


func TestBuildListOfTriangleNumbers(t *testing.T) {
	triangles := BuildListOfTriangleNumbers(10)
	expected := 55
	got := triangles[9]
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

