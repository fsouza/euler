package main

import (
	"container/vector"
	"testing"
)

func AssertGridProduct(t *testing.T, grid *vector.Vector, expected uint) {
	got := GridProduct(grid, 4)

	if expected != got {
		t.Errorf("The product for the grid \n\n%s\n\n should be %q, but was %q", ConvertGridToString(grid), expected, got)
	}
}

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

func TestGridProductOfHorizontalLine(t *testing.T) {
	firstLine := new(vector.IntVector)
	firstLine.Push(10)
	firstLine.Push(15)
	firstLine.Push(5)
	firstLine.Push(5)
	firstLine.Push(10)

	secondLine := new(vector.IntVector)
	secondLine.Push(8)
	secondLine.Push(5)
	secondLine.Push(18)
	secondLine.Push(9)
	secondLine.Push(105)

	grid := new(vector.Vector)
	grid.Push(firstLine)
	grid.Push(secondLine)

	for i := 0; i < 3; i++ {
		line := new(vector.IntVector)
		for j := 0; j < 5; j++ {
			line.Push(1)
		}
		grid.Push(line)
	}

	AssertGridProduct(t, grid, 85050)
}

func TestGridProductOfVerticalLine(t *testing.T) {
	grid := new(vector.Vector)

	for i := 0; i < 5; i++ {
		line := new(vector.IntVector)
		for j := 0; j < 5; j++ {
			if j == 0 {
				line.Push(10)
			} else {
				line.Push(1)
			}
		}

		grid.Push(line)
	}

	AssertGridProduct(t, grid, 10000)
}

func TestGridProductOfDiagonal(t *testing.T) {
	grid := new(vector.Vector)

	for i := 0; i < 5; i++ {
		line := new(vector.IntVector)
		for j := 0; j < 5; j++ {
			if i == j {
				line.Push(10)
			} else {
				line.Push(1)
			}
		}

		grid.Push(line)
	}

	AssertGridProduct(t, grid, 10000)
}
