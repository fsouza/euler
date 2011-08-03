package main

func CountLetterNumbers(word string) (sum int) {
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			sum += int(c - 'A' + 1)
		}
	}

	return
}

func TriangleNumberAt(position int) int {
	return position * (position + 1) / 2
}

func BuildListOfTriangleNumbers(limit int) []int {
	list := make([]int, limit)
	for i := 0; i < limit; i++ {
		list[i] = TriangleNumberAt(i + 1)
	}
	return list
}

func IsPresent(collection []int, item int) bool {
	for _, element := range collection {
		if item == element {
			return true
		}
	}

	return false
}

func IsTriangleWord(word string, triangleList []int) bool {
	wordValue := CountLetterNumbers(word)
	return IsPresent(triangleList, wordValue)
}
