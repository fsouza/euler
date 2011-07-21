package main

import (
	"sort"
	"strconv"
	"strings"
)

func HasPandigitalProduct(multiplier, multiplicand int) bool {
	sequence := []int{ '1', '2', '3', '4', '5', '6', '7', '8', '9', }
	strMultiplier := strconv.Itoa(multiplier)
	strMultiplicand := strconv.Itoa(multiplicand)

	if len(strMultiplier) < 3 && len(strMultiplicand) < 3 {
		return false
	}

	if len(strMultiplier) < 2 || len(strMultiplicand) < 2 {
		return false
	}

	both := strMultiplier + strMultiplicand
	for _, number := range both {
		if number == '0' || strings.Count(both, string(number)) > 1 {
			return false
		}
	}

	product := multiplier * multiplicand
	strProduct := strconv.Itoa(product)
	all := both + strProduct
	allBytes := []int(all)
	sort.Ints(allBytes)
	for i, number := range allBytes {
		if number != sequence[i] {
			return false
		}
	}

	return true
}
