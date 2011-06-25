package main

import (
	"strconv"
)

var unitsMap = map[int]int{
	1: 3,
	2: 3,
	3: 5,
	4: 4,
	5: 4,
	6: 3,
	7: 5,
	8: 5,
	9: 4,
	10: 3,
	11: 6,
	12: 6,
	20: 6,
}

// Only the differences :)
var leftMap = map[int]int {
	3:4,
	5:3,
}

func DivMod(divisor, numerator int) (div, mod int) {
	div = divisor / numerator
	mod = divisor % numerator
	return
}

func CountLetters(number int) int {
	if count, ok := unitsMap[number]; ok {
		return count
	}

	if number < 20 {
		_, mod := DivMod(number, 10)
		if count, ok := leftMap[mod]; ok {
			return count + 4
		}

		return unitsMap[mod] + 4
	}

	panic("I don't know how many letters " + strconv.Itoa(number) + " has.")
}
