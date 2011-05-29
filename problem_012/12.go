package main

import (
	"container/vector"
)

func GetPrimeFactors(number int) *vector.IntVector {
	numbers := new(vector.IntVector)

	primesChannel := GetPrimes()
	for  {
		prime := <-primesChannel

		if number < prime {
			break
		}

		if number % prime == 0 {
			numbers.Push(prime)
		}
	}

	return numbers
}

func main() {

}
