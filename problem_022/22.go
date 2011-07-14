package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func CountLetterNumbers(word string) (sum int64) {
	for _, c := range word {
		sum += int64(c - 'A' + 1)
	}

	return
}

func main() {
	totalScore := int64(0)

	if bytes, err := ioutil.ReadFile("names.txt"); err == nil {
		names := strings.Split(string(bytes), ",")
		sort.Strings(names)
		for i, name := range names {
			wordScore := CountLetterNumbers(name)
			wordScore *= int64(i + 1)
			totalScore += wordScore
		}
	}
	
	fmt.Println(totalScore)
}
