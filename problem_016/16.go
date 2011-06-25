package main

import (
	"strconv"
)

func SumDigits(number uint64) uint64 {
	var sum uint64

	digits := strconv.Uitoa64(number)
	for _, digit := range digits {
		sum += uint64(digit - '0')
	}

	return sum
}
