package main

import (
	"fmt"
	"strconv"
)

func SumDigits(number uint64) uint64 {
	var sum uint64

	digits := strconv.Uitoa64(number)
	for _, digit := range digits {
		sum += uint64(digit - '0')
	}

	return sum
}

func Pow(base, exponent uint64) uint64 {
	product := base

	for i := uint64(0); i < exponent - 1; i++ {
		product *= base
	}

	return product
}

func main() {
	// Not gonna work =/
	pow := Pow(2, 1000)
	fmt.Println(SumDigits(pow))
}
