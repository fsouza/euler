package main

import (
	"fmt"
)

func BuildSliceOfPrimes(limit int) []int {
	primesCh := GetPrimes(float64(limit))
	primes := make([]int, 0)

	for prime := range primesCh {
		if prime == 0 {
			break
		}

		primes = append(primes, prime)
	}

	return primes
}

func IsPrime(number int) bool {
	if number < 0 {
		number *= -1
	}

	if number < 4 {
		return true
	}

	if number % 2 == 0 {
		return false
	}

	count := 0
	for i := 5; i < number; i += 2 {
		if number % i == 0 {
			count++
		}
	}

	return count == 0
}

func main() {
	greater, result, primes := -1000, 0, BuildSliceOfPrimes(1000)
	for a := -999; a < 1000; a += 2 {
		for _, b := range primes {
			n := 1

			for IsPrime(n * n + a * n + b) {
				n++
			}

			if n > greater {
				greater, result = n, a * b
			}
		}
	}

	fmt.Println(result)
}
