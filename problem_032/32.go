package main

import (
	"sort"
	"strconv"
	"strings"
)

func IsPresent(numbers []int, value int) bool {
	for _, number := range numbers {
		if number == value {
			return true
		}
	}

	return false
}

func CanHavePandigitalProduct(multiplier, multiplicand int) bool {
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

	return true
}

func HasPandigitalProduct(multiplier, multiplicand int) (int, bool) {
	sequence := []int{ '1', '2', '3', '4', '5', '6', '7', '8', '9', }

	product := multiplier * multiplicand
	all := strconv.Itoa(multiplier) + strconv.Itoa(multiplicand) + strconv.Itoa(product)
	allBytes := []int(all)
	sort.Ints(allBytes)
	for i, number := range allBytes {
		if number != sequence[i] {
			return 0, false
		}
	}

	return product, true
}
