package main

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
