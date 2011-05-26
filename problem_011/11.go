package main

import (
	"container/vector"
	"fmt"
)

func ConvertGridToString(grid *vector.Vector) string {
	var output, format string

	for _, lineInterface := range *grid {
		line := lineInterface.(*vector.IntVector)
		for j, number := range *line {
			format = "%d"
			if j + 1 != line.Len() {
				format += " "
			}

			output += fmt.Sprintf(format, number)
		}

		if lineInterface != grid.Last() {
			output += "\n"
		}
	}

	return output
}

func GridProduct(grid *vector.Vector, elements int) uint {
	var greater, product uint

	for _, lineInterface := range *grid {
		line := lineInterface.(*vector.IntVector)
		for i := 0; i < line.Len() - elements + 1; i++ {
			product = 1
			for j := i; j < i + elements; j++ {
				product *= uint(line.At(j))
			}

			if product > greater {
				greater = product
			}
		}
	}

	return greater
}
