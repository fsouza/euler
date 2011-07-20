package main

import (
	"strconv"
)

func SumDigits(number int) (sum int) {
	digits := strconv.Itoa(number)

	for _, digit := range digits {
		sum += digit - '0'
	}

	return sum
}

func SumOfPowersOfDigits(exponent int) int {
	return 0
}
