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

func Pow(base, exponent uint64) uint64 {
	var product uint64 = 1

	for i := uint64(0); i < exponent; i++ {
		product *= base
	}

	return product
}
