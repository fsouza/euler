package main

import (
	"fmt"
	"strconv"
)

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

func (m Matrix) String() (out string) {
	for i, row := range m {
		for j, element := range row {
			out += strconv.Itoa(element)

			if j + 1 < len(row) {
				out += " "
			}
		}

		if i + 1 < len(m) {
			out += "\n"
		}
	}
	return
}

func (m Matrix) DiagonalsSum() int {
	middle := len(m) / 2
	sum := m[middle][middle]

	for i := 0; i < len(m); i++ {
		if i != middle {
			sum += m[i][i]
		}
	}

	for i, j := 0, len(m) - 1; i < len(m) && j >= 0; i, j = i + 1, j - 1 {
		if i != middle && j != middle {
			sum += m[i][j]
		}
	}

	return sum
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

func Right(matrix Matrix, currentRow, currentColumn, value int) (int, int) {
	currentColumn++
	matrix[currentRow][currentColumn] = value
	return currentRow, currentColumn
}

func Down(matrix Matrix, currentRow, currentColumn, value int) (int, int) {
	currentRow++
	matrix[currentRow][currentColumn] = value
	return currentRow, currentColumn
}

func Left(matrix Matrix, currentRow, currentColumn, value int) (int, int) {
	currentColumn--
	matrix[currentRow][currentColumn] = value
	return currentRow, currentColumn
}

func Top(matrix Matrix, currentRow, currentColumn, value int) (int, int) {
	currentRow--
	matrix[currentRow][currentColumn] = value
	return currentRow, currentColumn
}

type SetValue func(Matrix, int, int, int) (int, int)

func (s SetValue) Next() SetValue {
	if s == Right {
		return Down
	}

	if s == Down {
		return Left
	}

	if s == Left {
		return Top
	}

	return Right
}

func BuildSpiralMatrix(rows int) Matrix {
	var move SetValue

	total := rows * rows
	middle := rows / 2

	matrix := BuildMatrix(rows)
	matrix[middle][middle] = 1

	currentRow := middle
	currentColumn := middle
	movings := 1
	calls := 0

	move = Right
	i := 2
	for i <= total {
		for j := 0; j < movings; j++ {
			currentRow, currentColumn = move(matrix, currentRow, currentColumn, i)
			i++
		}

		move = move.Next()
		calls++

		if calls % 2 == 0 && currentRow != 0 {
			movings++
		}
	}

	return matrix
}

func main() {
	matrix := BuildSpiralMatrix(1001)
	fmt.Println(matrix.DiagonalsSum())
}
