package main

var unitsMap = map[int]int{
	0: 0,
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

func FindOnMaps(key int) int {
	if value, ok := leftMap[key]; ok {
		return value
	} else {
		return unitsMap[key]
	}

	panic("Shouldn't get here :(")
}

func DivMod(divisor, numerator int) (int, int) {
	return divisor / numerator, divisor % numerator
}

func CountUnits(number int) int {
	if number == 0 {
		return 0
	}

	return unitsMap[number]
}

func CountDozens(number int) int {
	if number == 0 {
		return 0
	}

	if number < 13 {
		return unitsMap[number]
	}

	var count int
	div, mod := DivMod(number, 10)

	find := true
	if div == 0 {
		find = false
	} else if div == 1 {
		count += 4
		count += FindOnMaps(mod)
		mod = 0
		find = false
	} else if div == 2 { // twenty
		count += unitsMap[20]
		find = false
	} else {
		count += 2
	}

	if find {
		count += FindOnMaps(div)
	}

	return count + CountUnits(mod)
}

func CountHundred(number int) int {
	if number == 0 {
		return 0
	}

	var count int
	div, mod := DivMod(number, 100)

	if div > 0 {
		count += 7 + unitsMap[div]
	}

	return count + CountDozens(mod)
}

func CountLetters(number int) int {
	return CountHundred(number)
}
