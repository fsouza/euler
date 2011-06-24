package main

import (
	"fmt"
)

func FindChainLength(number int) int {
	counter := 1
	for number != 1 {
		counter++
		if number % 2 == 0 {
			number /= 2
		} else {
			number = number * 3 + 1
		}
	}

	return counter
}

func main() {
	var length, greater, the_number int
	for i := 13; i <= 1000000; i++ {
		length = FindChainLength(i)
		if length > greater {
			greater = length
			the_number = i
		}
	}

	fmt.Println(greater)
	fmt.Println(the_number)
}
