package main

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
	var find, total int

	if count, ok := unitsMap[number]; ok {
		return count
	}

	div, mod := DivMod(number, 10)
	if div == 1 {
		total = 4
	} else if div == 2 { // damn twenty
		total = unitsMap[20]
	} else {
		total = 2
	}

	if mod == 0 {
		find = div
	} else {
		find = mod
	}

	if count, ok := leftMap[find]; ok {
		total += count
	} else {
		total += unitsMap[find]
	}

	return total
}
