package main

import (
	"sort"
	"strconv"
)

func IsPandigital(number int) bool {
	digits := []int(strconv.Itoa(number))
	sort.Ints(digits)

	sequence := make([]int, len(digits))
	for i, counter := digits[0], 0; i < digits[0] + len(digits); i, counter = i + 1, counter + 1 {
		sequence[counter] = i
	}

	for i, digit := range digits {
		if digit != sequence[i] {
			return false
		}
	}

	return true
}
