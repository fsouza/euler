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

func BuildLimitsForLength(number int) (minimun, maximun int) {
	minimun = 1
	for i := 1; i < number; i++ {
		minimun*= 10
	}

	maxString := strings.Repeat("9", number)
	maximun, _ = strconv.Atoi(maxString)

	return
}

func SumOfPowersOfDigits(exponent int) (sum int) {
	min, max := BuildLimitsForLength(exponent)

	for i := min; i <= max; i++ {
		if IsSumOfExponents(i, exponent) {
			sum += i
		}
	}

	return sum
}

