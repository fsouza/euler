package main

import (
	"fmt"
)

func main() {
	results := make(map[int]int)

	for a := 2; a <= 1000; a++ {
		for b := a; b <= 1000; b++ {
			for c := b; c <= 1000; c++ {
				if a + b + c < 1000 && a * a + b * b == c * c {
					if _, ok := results[a + b + c]; ok {
						results[a + b + c]++
					} else {
						results[a + b + c] = 1
					}
				}
			}
		}
	}

	greaterSum, greaterCount := 0, 0
	for sum, count := range results {
		if count > greaterCount {
			greaterSum, greaterCount = sum, count
		}
	}

	fmt.Println(greaterSum)
}
