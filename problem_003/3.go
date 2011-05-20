package main

func IsPrime(number int) bool {
	count := 0
	for i := 2; i < number; i++ {
		if number % i == 0 {
			count++
			break
		}
	}

	return count == 0
}

func Factor(i int) int {
	if i < 4 {
		return i
	}

	greater := i

	for j := greater; j >= 2; j-- {
		if IsPrime(j) && i % j == 0 {
			return j
		}
	}

	return greater
}
