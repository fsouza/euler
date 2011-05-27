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

func AreGridEquals(grid1, grid2 *vector.Vector) bool {
	if grid1.Len() != grid2.Len() {
		return false
	}

	for i := 0; i < grid1.Len(); i++ {
		line1 := grid1.At(i).(*vector.IntVector)
		line2 := grid2.At(i).(*vector.IntVector)
		for j := 0; j < line1.Len(); j++ {
			if line1.At(j) != line2.At(j) {
				return false
			}
		}
	}

	return true
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

func TestConvertGridFromString(t *testing.T) {
	grid := "10 15\n8 5"

	firstLine := new(vector.IntVector)
	firstLine.Push(10)
	firstLine.Push(15)

	secondLine := new(vector.IntVector)
	secondLine.Push(8)
	secondLine.Push(5)

	expected := new(vector.Vector)
	expected.Push(firstLine)
	expected.Push(secondLine)
	got := ReadGridFromString(grid)

	if !AreGridEquals(expected, got) {
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

func TestGridProductOfTheMainDiagonal(t *testing.T) {
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

func TestGridProductOfAnyTopDownDiagonal(t *testing.T) {
	stringGrid := "1 1 10 1 1 1\n1 1 1 10 1 1\n1 1 1 1 10 1\n1 1 1 1 1 10\n1 1 1 1 1 1"
	grid := ReadGridFromString(stringGrid)

	AssertGridProduct(t, grid, 10000)
}

func TestGridProductOfAnyDownTopDiagonal(t *testing.T) {
	stringGrid := "1 1 1 10 1\n1 1 10 1 1\n1 10 1 1 1\n10 1 1 1 1\n1 1 1 1 1"
	grid := ReadGridFromString(stringGrid)

	AssertGridProduct(t, grid, 10000)
}
