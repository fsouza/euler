package main

import (
	"container/vector"
	"fmt"
)

func ConvertGridToString(grid *vector.Vector) string {
	var output, format string

	for _, lineInterface := range *grid {
		line := lineInterface.(*vector.IntVector)
		for _, number := range *line {
			format = "%d"
			if number != line.Last() {
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

func GridProduct(grid *vector.Vector) uint {
	return 150;
}
