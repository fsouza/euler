package main

import (
	"strconv"
	"strings"
)

func BuildSequence(position int) string {
	if position < 10 {
		return strconv.Itoa(position)
	}

	var fractional string
	digits := make([]string, 0)

	for number := 1; len(fractional) < position; number++ {
		digits = append(digits, strconv.Itoa(number))
		fractional = strings.Join(digits, "")
	}

	return fractional
}

func FindNthDigit(position int) uint8 {
	fractional := BuildSequence(position)
	return fractional[position - 1] - '0'
}
