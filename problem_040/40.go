package main

import (
	"fmt"
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

func main() {
	sequence := BuildSequence(1000000)
	fmt.Println(
		(sequence[0] - '0') *
		(sequence[9] - '0') *
		(sequence[99] - '0') *
		(sequence[999] - '0') *
		(sequence[9999] - '0') *
		(sequence[99999] - '0') *
		(sequence[999999] - '0'))
}
