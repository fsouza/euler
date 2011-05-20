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
	return i
}

func main() {
}
