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

func main() {

}
