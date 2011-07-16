package main

import (
	"fmt"
	"strconv"
	"strings"
)

func permute(numbers []int) {
	length := len(numbers)
	key := length - 1

	for key > 0 && numbers[key] <= numbers[key - 1] {
		key--
	}

	key--

	if key < 0 {
		return
	}

	newKey := length - 1
	for newKey > 0 && numbers[newKey] <= numbers[key] {
		newKey--
	}

	numbers[key], numbers[newKey] = numbers[newKey], numbers[key]
	length--
	key++

	for length > key {
		numbers[length], numbers[key] = numbers[key], numbers[length]
		key++
		length--
	}
}

func FindNthPermutation(position int, numbers ...int) string {
	for counter := 0; counter < position; counter++ {
		permute(numbers)
	}

	strNumbers := make([]string, len(numbers))
	for i, number := range numbers {
		strNumbers[i] = strconv.Itoa(number)
	}

	return strings.Join(strNumbers, "")
}

func main() {
	fmt.Println(FindNthPermutation(999999, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
}
