package main

import (
	"strconv"
)

func BuildPrimesList(limit float64) []int {
	primes := make([]int, 0)
	primesCh := GetPrimes(limit)

	for prime := range primesCh {
		if prime == 0 {
			break
		}

		primes = append(primes, prime)
	}

	return primes
}

func IsPresent(slice []int, number int) bool {
	for _, element := range slice {
		if element == number {
			return true
		}
	}

	return false
}

func Truncate(number int, from string) int {
	digits := []int(strconv.Itoa(number))
	length := len(digits)

	if length == 1 {
		return 0
	}

	if from == "right" {
		digits = digits[:length - 1]
	} else {
		digits = digits[1:]
	}

	newNumber, _ := strconv.Atoi(string(digits))
	return newNumber
}

func TruncateRight(number int) int {
	return Truncate(number, "right")
}

func TruncateLeft(number int) int {
	return Truncate(number, "left")
}

func IsTruncatable(number int, primes []int) bool {
	operations := []string{"right", "left"}
	for _, operation := range operations {
		copyNumber := number
		for copyNumber = Truncate(copyNumber, operation); copyNumber != 0; copyNumber = TruncateRight(copyNumber) {
			if !IsPresent(primes, copyNumber) {
				return false
			}
		}
	}

	return true
}
