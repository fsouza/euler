package main

import (
	"strconv"
)

func HasPandigitalProduct(multiplier, multiplicand int) bool {
	strMultiplier := strconv.Itoa(multiplier)
	strMultiplicand := strconv.Itoa(multiplicand)

	if len(strMultiplier) < 3 && len(strMultiplicand) < 3 {
		return false
	}

	if len(strMultiplier) < 2 || len(strMultiplicand) < 2 {
		return false
	}

	return true
}
