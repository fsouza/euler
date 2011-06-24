package main

import (
	"strconv"
)

func SumDigitsOf(number uint) uint {
	digits := strconv.Uitoa(number)
	sum := uint(0)
	for _, digit := range digits {
		sum += uint(digit - '0')
	}

	return sum
}

func SliceDigits(number, length uint) string {
	digits := strconv.Uitoa(number)
	return digits[:length]
}

func Reduce(f func(x, y uint) uint, slice []uint) uint {
	result := slice[0]

	for _, value := range slice[1:len(slice)] {
		result = f(result, value)
	}

	return result
}
