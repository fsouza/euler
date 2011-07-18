package main

type Row []int
type Matrix []Row

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
