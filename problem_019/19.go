package main

import (
	"strconv"
)

func FindDayOfWeek(year, month, day int) int {
	var centuryCode int
	yearString := strconv.Itoa(year)
	lastYearDigits, _ := strconv.Atoi(yearString[2:])

	monthCodes := []int{ 1, 4, 4, 0, 2, 5, 0, 3, 6, 1, 4, 6 }

	if year < 2000 {
		centuryCode = 0
	} else {
		centuryCode = 6
	}

	yearCode := (centuryCode + lastYearDigits + (lastYearDigits / 4)) % 7
	return (yearCode + monthCodes[month - 1] + day) % 7
}
