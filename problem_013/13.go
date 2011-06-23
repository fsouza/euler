package main

import (
	"strconv"
)

func SumDigitsOf(number uint) uint {
	digits := strconv.Uitoa(number)
	sum := uint(0)
	for _, digit := range digits {
		sum += uint(digit - '0')
	}

	return sum
}
