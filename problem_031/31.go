package main

var coins = []int{ 1, 2, 5, 10, 20, 50, 100, }

func DivMod(divisor, numerator int) (int, int) {
	return divisor / numerator, divisor % numerator
}

func CalculateDifferentWays(value int) int {
	totalOfCoins := len(coins)
	return totalOfCoins
}
