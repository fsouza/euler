package main

import (
	"fmt"
)

func reverse(input string) string {
	chars := []int(input)
	length, half := len(chars), len(chars) / 2
	for i := 0; i < half; i++ {
		chars[i], chars[length - i - 1] = chars[length - i - 1], chars[i]
	}
	return string(chars)
}

func areStringEquals(str1, str2 string) bool {
	chars1 := []int(str1)
	chars2 := []int(str2)
	length := len(chars1)
	for i := 0; i < length; i++ {
		if chars1[i] != chars2[i] {
			return false
		}
	}
	return true
}

func main() {
	greater := 0

	for i := 999; i >= 900; i-- {
		for j := 999; j >= 900; j-- {
			result := i * j
			resultString := fmt.Sprintf("%d", result)
			reversed := reverse(resultString)
			if (areStringEquals(resultString, reversed) && result > greater) {
				greater = result
			}
		}
	}

	fmt.Println(greater)
}
