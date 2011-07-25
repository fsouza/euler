package main

import (
	"sort"
)

func IsPandigital(input string) bool {
	sequence := []int{ '1', '2', '3', '4', '5', '6', '7', '8', '9', }
	digits := []int(input)
	sort.Ints(digits)

	if len(digits) == len(sequence) {
		for i, digit := range digits {
			if digit != sequence[i] {
				return false
			}
		}

		return true
	}

	return false
}
