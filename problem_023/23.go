package main

import (
	"math"
)

// Stole from problem 21
func SumOfFactors(number uint) uint {
	sum := uint(1)
	sqrt := uint(math.Sqrt(float64(number)))

	for i := uint(2); i < sqrt + 1; i++ {
		if number % i == 0 {
			sum += i + number / i
		}
	}

	return sum
}

func IsAbundant(number uint) bool {
	return SumOfFactors(number) > number
}
