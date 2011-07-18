package main

type Row []int
type Matrix []Row

func (m Matrix) Equals(other Matrix) bool {
	if len(m) != len(other) {
		return false
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] != other[i][j] {
				return false
			}
		}
	}

	return true
}

func BuildMatrix(rows int) Matrix {
	if rows % 2 == 0 {
		panic("I can't build matrixes with even lines")
	}

	matrix := make(Matrix, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make(Row, rows)
	}

	return matrix
}

var BuildSpiralMatrix = BuildMatrix
