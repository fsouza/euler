package main

import (
	"fmt"
)

func PowDifference(number int) int {
	sumOfPows := 0
	sumToPow := 0
	for i := 1; i <= number; i++ {
		sumToPow += i
		sumOfPows += i * i
	}
	powOfSum := sumToPow * sumToPow
	return powOfSum - sumOfPows
}

func main() {
	fmt.Println(PowDifference(100))
}
