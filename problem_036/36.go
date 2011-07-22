package main

func Reverse(slice []int) []int {
	length, half := len(slice), len(slice) / 2

	for i := 0; i < half; i++ {
		slice[i], slice[length - i - 1] = slice[length - i - 1], slice[i]
	}

	return slice
}

func DivMod(numerator, denominator int) (int, int) {
	return numerator / denominator, numerator % denominator
}

func ToBinary(number int) string {
	digits := []int{}

	div, mod := DivMod(number, 2)
	digits = append(digits, mod + '0')
	for div > 1 {
		div, mod = DivMod(div, 2)
		digits = append(digits, mod + '0')
	}
	digits = append(digits, div + '0')

	return string(Reverse(digits))
}

func IsTheSameReversed(s string) bool {
	digits := []int(s)
	reversed := Reverse(digits)
	return s == string(reversed)
}
