package main

import (
	"container/vector"
	"fmt"
	"math"
)

func GetPrimeFactors(number int) (*vector.IntVector, *vector.IntVector) {
	numbers := new(vector.IntVector)
	counts := new(vector.IntVector)

	primesChannel := GetPrimes()
	result := number
	for  {
		prime := <-primesChannel

		if number < prime {
			break
		}

		if result % prime == 0 {
			numbers.Push(prime)
			count := 1

			for result = result / prime; result % prime == 0; result = result / prime {
				count++
			}

			counts.Push(count)
		}
	}

	return numbers, counts
}

func CountDivisors(number int) int {
	_, factorsCounts := GetPrimeFactors(number)
	count := 1

	for i := 0; i < factorsCounts.Len(); i++ {
		count *= (factorsCounts.At(i) + 1)
	}

	return count
}

func FindTriangleNumberAt(index int, cache map[int] int) int {
	_, foundOnCache := cache[index]
	if !foundOnCache {
		if index == 1 {
			return index
		}

		cache[index] = index + FindTriangleNumberAt(index - 1, cache)
	}

	return cache[index]
}

func FindFirstNumberToOverLimit(limit int) int {
	triangleCache := make(map[int] int)

	minimum := int(math.Pow(float64(limit), 2))
	for i := 1; ; i++ {
		triangle := FindTriangleNumberAt(i, triangleCache)

		if triangle > minimum {
			divisorsCount := CountDivisors(triangle)

			if divisorsCount > limit {
				return triangle
			}
		}
	}

	return 0
}

func main() {
	fmt.Println(FindFirstNumberToOverLimit(50))
}
