package main

import (
	"testing"
)

func TestTriangleElementHasValue(t *testing.T) {
	element := Element{68}
	if element.value != 68 {
		t.Errorf("Element should have a value!")
	}
}

func TestTriangleElementWithAdjacents(t *testing.T) {
	element := NewElement(68, nil, nil)
	if element.value != 68 || element.right != nil || element.left != nil {
		t.Errorf("Element with adjacents should have a value and two adjacents")
	}
}

