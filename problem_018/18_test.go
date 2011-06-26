package main

import (
	"testing"
)

func AssertEqual(expected, got int, t *testing.T) {
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

func TestGetNewElement(t *testing.T) {
	element := NewElement(100, nil, nil)
	if element.value != 100 {
		t.Errorf("Ops! something went wrong and I didn't set the value properly")
	}
}

func TestGetNewElementWithAdjacents(t *testing.T) {
	left := NewElement(50, nil, nil)
	right := NewElement(100, nil, nil)
	element := NewElement(150, left, right)
	if element.value != 150 || element.left.value != 50 || element.right.value != 100 {
		t.Errorf("Something went wrong setting value for adjacents! :(")
	}
}

func TestBuildStructureFromString(t *testing.T) {
	theString := "3\n7 4"
	root := BuildFromString(theString)
	AssertEqual(3, root.value, t)
	AssertEqual(7, root.left.value, t)
	AssertEqual(4, root.right.value, t)
	// AssertEqual(2, root.left.left.value, t)
	// AssertEqual(4, root.left.right.value, t)
	// AssertEqual(4, root.right.left.value, t)
	// AssertEqual(6, root.right.right.value, t)
}

