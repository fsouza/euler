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
}

var 

func CountLetters(number int) int {
	if number <= 12 {
		return unitsMap[number]
	}

	panic("I don't know how many letters this number has.")
}
