package main

import (
	"fmt"
	"strconv"
)

func IsPresent(slice []int, number int) bool {
	for _, element := range slice {
		if element == number {
			return true
		}
	}

	return false
}

func Rotate(s string ) string {
	digits := []int(s)
	length := len(digits)

	the_digit := digits[0]
	for i := length - 1; i >= 0; i-- {
		the_digit, digits[i] = digits[i], the_digit
	}

	return string(digits)
}

func IsCircular(number int, primes []int) bool {
	strNumber := strconv.Itoa(number)
	rotated := Rotate(strNumber)
	for rotated != strNumber {
		rotatedNumber, _ := strconv.Atoi(rotated)
		if !IsPresent(primes, rotatedNumber) {
			return false
		}
		rotated = Rotate(rotated)
	}

	return true
}

func BuildPrimesList(limit float64) []int {
	primes := make([]int, 0)
	primesCh := GetPrimes(limit)
	
	for prime := range primesCh {
		if prime == 0 {
			break
		}

		primes = append(primes, prime)
	}

	return primes
}

func CountCircularPrimes(limit float64) int {
	count := 0
	primes := BuildPrimesList(limit)

	for _, prime := range primes {
		if IsCircular(prime, primes) {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(CountCircularPrimes(1000000))
}
