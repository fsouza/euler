package main

import (
	"container/vector"
)

func GetPrimeFactors(number int) (*vector.IntVector, *vector.IntVector) {
	numbers := new(vector.IntVector)
	counts := new(vector.IntVector)

	primesChannel := GetPrimes()
	result := number
	for  {
		prime := <-primesChannel

		if result % prime == 0 {
			numbers.Push(prime)
			count := 1

			for result = result / prime; result % prime == 0; result = result / prime {
				count++
			}

			counts.Push(count)
		}

		if number < prime {
			break
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

func FindTriangleNumberAt(index int) int {
	if index == 1 {
		return index
	}

	return index + FindTriangleNumberAt(index - 1)
}

func main() {

}
