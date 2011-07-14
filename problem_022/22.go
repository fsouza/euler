package main

func CountLetterNumbers(word string) (sum int) {
	for _, c := range word {
		sum += c - 64
	}

	return
}
