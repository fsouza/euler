package main

import (
	"fmt"
)

func IsPrime(number uint64) bool {
	count := 0
	for i := uint64(2); i < number; i++ {
		if number % i == 0 {
			count++
			break
		}
	}

	return count == 0
}

func Factor(i uint64) uint64 {
	if i < 4 {
		return i
	}

	greater := i

	for j := greater; j >= 2; j-- {
		if IsPrime(j) && i % j == 0 {
			return j
		}
	}

	return greater
}

func main() {
	fmt.Println(Factor(600851475143))
}
