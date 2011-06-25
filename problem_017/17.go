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
	8:4,
}

func DivMod(divisor, numerator int) (div, mod int) {
	div = divisor / numerator
	mod = divisor % numerator
	return
}

func CountLetters(number int) int {
	var find, total int

	if number == 0 {
		return 0
	}

	if count, ok := unitsMap[number]; ok {
		return count
	}

	div, mod := DivMod(number, 10)
	if div == 1 {
		total = 4
		find = mod
	} else if div == 2 { // damn twenty
		total = unitsMap[20]
		find = mod
	} else {
		total = 2
		find = div
	}

	if count, ok := leftMap[find]; ok {
		total += count
	} else {
		total += unitsMap[find]
	}

	if find == mod {
		return total
	}

	return total + CountLetters(mod)
}
