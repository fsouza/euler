package main

import (
	"big"
)

func Fibonacci(number int64) *big.Int {
	first, second := big.NewInt(0), big.NewInt(1)

	for i := number; i > 0; i-- {
		first, second = second, first.Add(first, second)
	}

	return first
}
