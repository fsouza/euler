package main

import (
	"fmt"
	"sort"
	"strconv"
)

func IsPandigital(number int) bool {
	digits := []int(strconv.Itoa(number))
	sort.Ints(digits)

	sequence := make([]int, len(digits))
	for i, counter := '1', 0; i < len(digits) + '1'; i, counter = i + 1, counter + 1 {
		sequence[counter] = i
	}

	for i, digit := range digits {
		if digit != sequence[i] {
			return false
		}
	}

	return true
}

func main() {
	max := 0
	primes := GetPrimes(10000000)
	for prime := range primes {
		if prime == 0 {
			break
		}

		if IsPandigital(prime) && prime > max {
			max = prime
		}
	}

	fmt.Println(max)
}
