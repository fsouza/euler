package main

var coins = []int{ 1, 2, 5, 10, 20, 50, 100, }

func DivMod(divisor, numerator int) (int, int) {
	return divisor / numerator, divisor % numerator
}

func CalculateDifferentWays(value int) (total int) {
	totalOfCoins := len(coins)

	for i := totalOfCoins - 1; i >= 0; i-- {
		div := value / coins[i]
		if div > 0 {
			total++
		}
	}

	return
}