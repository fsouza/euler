package main

import (
	"strconv"
	"strings"
)

func IntPow(base, exponent int) int {
	result := base

	for i := 1; i < exponent; i++ {
		result *= base
	}

	return result
}

func IsSumOfExponents(number, exponent int) bool {
	digits := strconv.Itoa(number)
	sum := 0

	for _, digit := range digits {
		sum += IntPow(digit - '0', exponent)
	}

	return sum == number
}

func LenOfNumber(number int) int {
	digits := strconv.Itoa(number)
	return len(digits)
}

func SumOfPowersOfDigits(exponent int) (sum int) {
	return sum
}

func BuildLimitsForLength(number int) (minimun, maximun int) {
	minimun = 1
	for i := 1; i < number; i++ {
		minimun*= 10
	}

	maxString := strings.Repeat("9", number)
	maximun, _ = strconv.Atoi(maxString)

	return
}
