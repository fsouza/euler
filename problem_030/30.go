package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func IntPow(base, exponent int) int {
	result := base

	for i := 1; i < exponent; i++ {
		result *= base
	}

	return result
}

func IsSumOfExponents(number, exponent int) bool {
	digits := strconv.Itoa(number)
	sum := 0

	for _, digit := range digits {
		sum += IntPow(digit - '0', exponent)
	}

	return sum == number
}

func sumPowserFromTo(exponent, from, to int, ch chan<- int) {
	sum := 0

	for i := from; i < to; i++ {
		if IsSumOfExponents(i, exponent) {
			sum += i
		}
	}

	ch <- sum
}

func getMinAndMax(perIteration, remainder, currentIteration int) (min, max int) {
	if currentIteration == 0 {
		min = 2
		max = perIteration + remainder
	} else {
		min = perIteration + remainder + perIteration * (currentIteration - 1)
		max = perIteration + remainder + perIteration * currentIteration
	}

	return
}

func SumOfPowersOfDigits(exponent int) (sum int) {
	const (
		CPUS = 4
		MAX = 9999999
	)
	oldMaxProcs := runtime.GOMAXPROCS(CPUS)

	perIteration := MAX / CPUS
	remainder := MAX % CPUS

	ch := make(chan int, CPUS)
	for i := 0; i < CPUS; i++ {
		min, max := getMinAndMax(perIteration, remainder, i)
		go sumPowserFromTo(exponent, min, max, ch)
	}

	for i := 0; i < CPUS; i++ {
		s := <-ch
		sum += s
	}

	runtime.GOMAXPROCS(oldMaxProcs)

	return sum
}

func main() {
	fmt.Println(SumOfPowersOfDigits(5))
}
