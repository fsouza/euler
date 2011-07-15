package main

import (
	"fmt"
	"math"
)

// Stole from problem 21
func SumOfFactors(number uint) uint {
	sum := uint(1)
	sqrt := uint(math.Sqrt(float64(number)))

	for i := uint(2); i < sqrt + 1; i++ {
		if number % i == 0 {
			sum += i

			result := number / i
			if i != result {
				sum += result
			}
		}
	}

	return sum
}

func IsAbundant(number uint) bool {
	return SumOfFactors(number) > number
}

func main() {
	var abundantCount int

	limit := uint(20162)
	abundants := make([]uint, 500)
	abundantSums := make(map[uint]bool)

	for i := uint(1); i < limit; i++ {
		if IsAbundant(i) {
			if abundantCount < 500 {
				abundants[abundantCount] = i
			} else {
				abundants = append(abundants, i)
			}

			abundantCount++
		}
	}

	for i := 0; i < abundantCount; i++ {
		for j := i; j < abundantCount; j++ {
			sum := abundants[i] + abundants[j]
			if sum < limit {
				abundantSums[sum] = true
			}
		}
	}

	sum := uint(0)
	for i := uint(1); i < limit; i++ {
		if _, ok := abundantSums[i]; !ok {
			sum += i
		}
	}

	fmt.Println(sum)
}
