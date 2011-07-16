package main

import (
	"big"
	"fmt"
)

func Fibonacci(number int64) *big.Int {
	first, second := big.NewInt(0), big.NewInt(1)

	for i := number; i > 0; i-- {
		first, second = second, first.Add(first, second)
	}

	return first
}

func main() {
	number := int64(0)
	for len(Fibonacci(number).String()) < 1000 { number++ }
	fmt.Println(number)
}
