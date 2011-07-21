package main

import (
	"strconv"
)

func IsNonTrivial(numerator, denominator int) bool {
	if numerator == denominator {
		return false
	}

	strNumerator := strconv.Itoa(numerator)
	strDenominator := strconv.Itoa(denominator)

	if len(strNumerator) != 2 || len(strDenominator) != 2 {
		return false
	}

	if strNumerator[0] == strDenominator[1] {
		return float64(numerator) / float64(denominator) == float64(strNumerator[1] - '0') / float64(strDenominator[0] - '0')
	}

	if strNumerator[1] == strDenominator[0] {
		return float64(numerator) / float64(denominator) == float64(strNumerator[0] - '0') / float64(strDenominator[1] - '0')
	}

	return false
}
