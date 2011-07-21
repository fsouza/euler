package main

import (
	"big"
)

// Function imported from problem 20 (where it's tested)
func Factorial(number int64) *big.Int {
	factorial := big.NewInt(1)

	for i := number; i >= 1; i-- {
		factorial.Mul(factorial, big.NewInt(i))
	}

	return factorial
}
