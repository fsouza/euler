package main

import (
	"fmt"
)

func GetPascalsTriangleCell(row, column int) uint64 {
	entireRow := make([]uint64, row + 1)
	entireRow[0] = 1

	numerator := uint64(row)
	denominator := uint64(1)

	for counter := 1; counter < row + 1; counter++ {
		entireRow[counter] = entireRow[counter - 1] * numerator/denominator
		numerator--
		denominator++
	}

	return entireRow[column]
}

func main() {
	fmt.Println(GetPascalsTriangleCell(20 * 2, 20))
}
