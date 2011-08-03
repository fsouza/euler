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
