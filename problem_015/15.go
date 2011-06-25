package main

type PascalRow []int

func GetPascalsTriangleCell(row, column int) int {
	entireRow := make([]int, row + 1)
	entireRow[0] = 1

	numerator := row
	denominator := 1

	for counter := 1; counter < row + 1; counter++ {
		entireRow[counter] = entireRow[counter - 1] * numerator/denominator
		numerator--
		denominator++
	}

	return entireRow[column]
}
