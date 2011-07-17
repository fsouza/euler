package main

type Slice []int

func (s Slice) IsElementPresent(number int) bool {
	for _, element := range []int(s) {
		if number == element {
			return true
		}
	}

	return false
}

func prepareNumerator(numerator, denominator int) int {
	for numerator < denominator {
		numerator *= 10
	}

	return numerator
}

func FindLongestRecurring(limit int) (number int) {
	greater := 0

	for i := 2; i < limit; i++ {
		visited := make(Slice, 0)
		numerator := prepareNumerator(1, i)

		done := false
		remainders := 0
		for !done {
			numerator = numerator % i

			if numerator == 0 || visited.IsElementPresent(numerator) {
				done = true
			} else {
				visited = append(visited, numerator)
				numerator = prepareNumerator(numerator, i)
			}

			remainders++
		}

		if remainders > greater {
			greater = remainders
			number = i
		}
	}
	return
}
