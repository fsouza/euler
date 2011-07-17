package main

func BuildSliceOfPrimes(limit int) []uint {
	primesCh := GetPrimes(float64(limit))
	primes := make([]uint, 0)

	for prime := range primesCh {
		if prime == 0 {
			break
		}

		primes = append(primes, prime)
	}

	return primes
}
