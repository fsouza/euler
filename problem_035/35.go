package main

import (
	"strconv"
)

func IsPresent(slice []int, number int) bool {
	for _, element := range slice {
		if element == number {
			return true
		}
	}

	return false
}

func Rotate(number int) int {
	numberDigits := []int(strconv.Itoa(number))
	length := len(numberDigits)

	the_digit := numberDigits[0]
	for i := length - 1; i >= 0; i-- {
		the_digit, numberDigits[i] = numberDigits[i], the_digit
	}

	newNumber, _ := strconv.Atoi(string(numberDigits))
	return newNumber
}

func IsCircular(number int, primes []int) bool {
	rotated := Rotate(number)
	for rotated != number {
		if !IsPresent(primes, rotated) {
			return false
		}
		rotated = Rotate(rotated)
	}

	return true
}

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
