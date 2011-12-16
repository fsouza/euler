package main

import (
	"fmt"
	"math/big"
	"strconv"
)

// Function imported from problem 20 (where it's tested)
func Factorial(number int64) *big.Int {
	factorial := big.NewInt(1)

	for i := number; i >= 1; i-- {
		factorial.Mul(factorial, big.NewInt(i))
	}

	return factorial
}

func main() {
	factorials := make([]*big.Int, 0)
	for i := int64(0); i < 10; i++ {
		factorials = append(factorials, Factorial(i))
	}

	sum := int64(0)
	for i := int64(10); i < 50000; i++ {
		factorialSum := int64(0)
		for _, digit := range strconv.FormatInt(i, 10) {
			factorialSum += factorials[digit-'0'].Int64()
		}

		if factorialSum == i {
			sum += i
		}
	}

	fmt.Println(sum)
}
