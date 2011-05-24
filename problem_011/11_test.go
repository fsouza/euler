package main

import (
	"container/vector"
	"testing"
)

func TestConvertGridToString(t *testing.T) {
	firstLine := new(vector.IntVector)
	firstLine.Push(10)
	firstLine.Push(15)

	secondLine := new(vector.IntVector)
	secondLine.Push(8)
	secondLine.Push(5)

	grid := new(vector.Vector)
	grid.Push(firstLine)
	grid.Push(secondLine)

	expected := "10 15\n8 5"
	got := ConvertGridToString(grid)

	if expected != got {
		t.Errorf("Expected: %q\nGot: %q", expected, got)
	}
}

func TestProductOfSimpleGrid(t *testing.T) {
	firstLine := new(vector.IntVector)
	firstLine.Push(10)
	firstLine.Push(15)

	secondLine := new(vector.IntVector)
	secondLine.Push(8)
	secondLine.Push(5)

	grid := new(vector.Vector)
	grid.Push(firstLine)
	grid.Push(secondLine)

	expected := uint(150)
	got := GridProduct(grid)

	if expected != got {
		t.Errorf("The product for the grid \n\n%s\n\n should be %q, but was %q", ConvertGridToString(grid), expected, got)
	}
}
