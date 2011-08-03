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

func TestCanFindTheSumOfAWord(t *testing.T) {
	expected := 53
	got := CountLetterNumbers("COLIN")
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

func TestIstriangleWord(t *testing.T) {
	triangleList := BuildListOfTriangleNumbers(20)
	input := "SKY"
	expected := true
	got := IsTriangleWord(input, triangleList)
	if expected != got {
		t.Errorf("Assertion error.\nExpected %v to be a triangle word? %v.\nGot: %v.", input, expected, got)
	}
}

func TestIsPresent(t *testing.T) {
	expected := true
	got := IsPresent([]int{1, 2, 3}, 2)
	if expected != got {
		t.Errorf("Assertion error.\nExpected %v to be present? %v.\nGot: %v.", 2, expected, got)
	}
}

