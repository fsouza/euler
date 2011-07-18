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

