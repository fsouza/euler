package main

import (
	"big"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func CountLetterNumbers(word string) *big.Int {
	sum := big.NewInt(0)

	for _, c := range word {
		sum.Add(sum, big.NewInt(int64(c - 'A' + 1)))
	}

	return sum
}

func main() {
	totalScore := big.NewInt(0)

	if bytes, err := ioutil.ReadFile("names.txt"); err == nil {
		names := strings.Split(string(bytes), ",")
		sort.Strings(names)
		for i, name := range names {
			wordScore := CountLetterNumbers(name)
			wordScore.Mul(wordScore, big.NewInt(int64(i + 1)))
			totalScore.Add(totalScore, wordScore)
		}
	}
	
	fmt.Println(totalScore)
}
