package main

import (
	"big"
	"fmt"
)

func SumDigitsOf(number *big.Int) int {
	digits := number.String()
	sum := int(0)
	for _, digit := range digits {
		sum += int(digit - '0')
	}

	return sum
}

func Factorial(number int64) *big.Int {
	factorial := big.NewInt(1)

	for i := number; i >= 1; i-- {
		factorial.Mul(factorial, big.NewInt(i))
	}

	return factorial
}

func main() {
	fact := Factorial(100)
	fmt.Println(SumDigitsOf(fact))
}
