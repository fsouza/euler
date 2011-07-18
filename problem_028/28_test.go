package main

import (
	"testing"
)

func TestRowType(t *testing.T) {
	row := make(Row, 10)
	if row[0] != 0 {
		t.Errorf("Row type missing!")
	}
}


func TestMatrixType(t *testing.T) {
	matrix := make(Matrix, 10)
	if matrix[0] != nil {
		t.Errorf("Matrix type missing!")
	}
}

func TestBuildMatrix(t *testing.T) {
	matrix := BuildMatrix(3)
	if matrix[0][0] != 0 {
		t.Errorf("The matrix was built wrong")
	}
}


func TestCanBuildMatrixOnlyForOddRowsAndCols(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Should fail trying to create a matrix with even number of columns/rows")
		}
	}()
	BuildMatrix(4)
}

func TestCheckIfMatrixesEquals(t *testing.T) {
	matrix1 := Matrix {
		Row { 1, 2, 3, },
		Row { 4, 5, 6, },
		Row { 7, 8, 9, },
	}

	matrix2 := Matrix {
		Row { 1, 2, 3, },
		Row { 4, 5, 6, },
		Row { 7, 8, 9, },
	}

	if !matrix1.Equals(matrix2) {
		t.Errorf("The Equals method should consider the two matrixes equal.")
	}
}

func TestMatrixAsString(t *testing.T) {
	matrix := Matrix {
		Row { 1, 2, 3, },
		Row { 4, 5, 6, },
		Row { 7, 8, 9, },
	}

	expected := "1 2 3\n4 5 6\n7 8 9"
	got := matrix.String()
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}


func TestBuildSpiralMatrix(t *testing.T) {
	expected := Matrix {
		Row { 21, 22, 23, 24, 25, },
		Row { 20, 7, 8, 9, 10, },
		Row { 19, 6, 1, 2, 11, },
		Row { 18, 5, 4, 3, 12, },
		Row { 17, 16, 15, 14, 13 },
	}

	got := BuildSpiralMatrix(5)
	if !expected.Equals(got) {
		t.Errorf("Assertion error.\nExpected:\n%v\n\n\nGot:\n%v", expected, got)
	}
}

func TestSumMatrixDiagonals(t *testing.T) {
	matrix := BuildSpiralMatrix(5)

	expected := 101
	got := matrix.DiagonalsSum()
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

