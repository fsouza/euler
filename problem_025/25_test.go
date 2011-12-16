package main

import (
	"math/big"
	"testing"
)

type FibonacciTest struct {
	in  int64
	out *big.Int
}

var FibonacciTests = []FibonacciTest{
	FibonacciTest{3, big.NewInt(2)},
	FibonacciTest{7, big.NewInt(13)},
	FibonacciTest{9, big.NewInt(34)},
}

func TestFibonacci(t *testing.T) {
	for _, test := range FibonacciTests {
		got := Fibonacci(test.in)
		if test.out.Cmp(got) != 0 {
			t.Errorf("Assertion error. Expected: %v\nGot: %v", test.out, got)
		}
	}
}
