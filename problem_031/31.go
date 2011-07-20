package main

var coins = []int{ 1, 2, 5, 10, 20, 50, 100, 200, }

func DivMod(divisor, numerator int) (int, int) {
	return divisor / numerator, divisor % numerator
}

func calculate(value, coinPosition int) int {
	if value == 0 {
		return 1
	}

	if value < 0 || coinPosition < 0 {
		return 0
	}

	return calculate(value, coinPosition - 1) + calculate(value - coins[coinPosition], coinPosition)
}

func CalculateDifferentWays(value int) int {
	i := 0

	for i = len(coins) - 1; i >= 0; i-- {
		if value / coins[i] > 0 {
			break
		}
	}

	return calculate(value, i)
}
