package main

import (
	"strconv"
	"strings"
)

func FindNthDigit(position int) uint8 {
	var fractional string
	digits := make([]string, 0)

	for number := 1; len(fractional) < position; number++ {
		digits = append(digits, strconv.Itoa(number))
		fractional = strings.Join(digits, "")
	}

	return fractional[position] - '0'
}
