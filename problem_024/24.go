package main

import (
	"sort"
)

func Factorial(number int) int {
	factorial := int(1)

	for i := number; i >= 1; i-- {
		factorial *= i
	}

	return factorial
}

func GeneratePermutations(numbers ...int) []string {
	permutations := make([]string, Factorial(len(numbers)))

	sort.Strings(permutations)
	return permutations
}
