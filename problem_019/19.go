package main

import (
	"fmt"
	"strconv"
)

func FindDayOfWeek(year, month, day int) (weekday int) {
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
	weekday = (yearCode + monthCodes[month - 1] + day) % 7

	if (year % 400 == 0 || (year % 100 != 0 && year % 4 == 0)) && month < 3 {
		weekday--
	}

	return weekday
}

func main() {
	var sundayCounter int

	for currentYear := 1901; currentYear < 2001; currentYear++ {
		for month := 1; month < 13; month++ {
			weekday := FindDayOfWeek(currentYear, month, 1)
			if weekday == 1 {
				sundayCounter++
			}
		}
	}

	fmt.Println(sundayCounter)
}
