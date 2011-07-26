package main

import (
	"strconv"
)

func BuildSequence(limit int) string {
	if limit < 10 {
		return strconv.Itoa(limit)
	}
	var fractional string

	for number := 1; len(fractional) < limit; number++ {
		fractional += strconv.Itoa(number)
	}

	return fractional
}

func FindNthDigit(position int) uint8 {
	fractional := BuildSequence(position)
	return fractional[position - 1] - '0'
}
