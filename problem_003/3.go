package main

import (
	"fmt"
	"math"
)

func IsPrime(number int) bool {
	count := 0
	for i := 2; i < number; i++ {
		if number % i == 0 {
			count++
			break
		}
	}

	return count == 0
}

func Factor(i int) int {
	if i < 4 {
		return i
	}

	greater := int(math.Sqrt(float64(i)))
	fmt.Println(greater)
	return greater
}

func main() {
}
