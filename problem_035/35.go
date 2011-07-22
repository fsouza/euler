package main

import (
	"strconv"
)

func IsPrime(number int) bool {
	if number < 0 {
		number *= -1
	}

	if number < 4 {
		return true
	}

	if number % 2 == 0 {
		return false
	}

	count := 0
	for i := 5; i < number; i += 2 {
		if number % i == 0 {
			count++
		}
	}

	return count == 0
}

func Rotate(number int) int {
	numberDigits := []int(strconv.Itoa(number))
	length := len(numberDigits)

	the_digit := numberDigits[0]
	for i := length - 1; i >= 0; i-- {
		the_digit, numberDigits[i] = numberDigits[i], the_digit
	}

	newNumber, _ := strconv.Atoi(string(numberDigits))
	return newNumber
}

func IsCircular(number int) bool {
	rotated := Rotate(number)
	for rotated != number {
		if !IsPrime(rotated) {
			return false
		}
		rotated = Rotate(rotated)
	}

	return true
}
