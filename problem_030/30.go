package main

import (
	"fmt"
	"strconv"
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

func SumOfPowersOfDigits(exponent int) (sum int) {
	const MAX = 9999999

	for i := 2; i <= MAX; i++ {
		if IsSumOfExponents(i, exponent) {
			sum += i
		}
	}

	return sum
}

func main() {
	fmt.Println(SumOfPowersOfDigits(5))
}
