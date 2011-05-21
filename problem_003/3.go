package main

import (
	"fmt"
)

func Factor(value uint64) uint64 {
	factor := uint64(3)

	for factor < value {
		if value % factor == 0 {
			value /= factor
		} else {
			factor += 2
		}
	}

	return factor
}

func main() {
	fmt.Println(Factor(600851475143))
}
