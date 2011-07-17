package main

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
