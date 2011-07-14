package main

import (
	"fmt"
	"math"
)

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

func main() {
	sum := uint(0)

	for i := uint(2); i < 10000; i++ {
		sumOfI := SumOfFactors(i)
		if sumOfI > i && SumOfFactors(sumOfI) == i {
			sum += sumOfI + i
		}
	}

	fmt.Println(sum)
}
