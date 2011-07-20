package main

import (
	"big"
	"fmt"
)

func IsPresent(sequence []*big.Int, number *big.Int) bool {
	for _, element := range sequence {
		if element.Cmp(number) == 0 {
			return true
		}
	}

	return false
}

func AddToSequence(sequence []*big.Int, numbers ...*big.Int) []*big.Int {
	for _, number := range numbers {
		if !IsPresent(sequence, number) {
			sequence = append(sequence, number)
		}
	}

	return sequence
}

func CountTermsInSequence(a, b int64) int {
	sequence := make([]*big.Int, 0)
	base := big.NewInt(a)
	exponent := big.NewInt(b)

	for i := big.NewInt(2); i.Cmp(base) < 1; i.Add(i, big.NewInt(1)) {
		for j := big.NewInt(2); j.Cmp(exponent) < 1; j.Add(j, big.NewInt(1)) {
			sequence = AddToSequence(sequence, big.NewInt(0).Exp(i, j, nil))
		}
	}

	return len(sequence)
}

func main() {
	fmt.Println(CountTermsInSequence(100, 100))
}
