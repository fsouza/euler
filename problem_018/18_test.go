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

