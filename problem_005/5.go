package main

import (
	"fmt"
)

func main() {
	numbersToCheck := []int{19, 18, 17, 16, 15, 14, 13, 12, 11}
	total := len(numbersToCheck)
	found := false
	for i := 2520; found == false; i+=20 {
		found = true
		for j := 0; j < total; j++ {
			if i % numbersToCheck[j] != 0 {
				found = false
			}
		}

		if found {
			fmt.Println(i)
		}
	}
}
