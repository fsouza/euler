package main

import (
	"testing"
)

func TestIfTheMultiplicandAndMultiplierHasLessThan5DigitsTogetherTheProductIsntPanDigital(t *testing.T) {
	type SmallMultiplierTest struct {
		multiplier, multiplicand int
	}

	SmallMultiplierTests := []SmallMultiplierTest{
		SmallMultiplierTest { 2, 4, },
		SmallMultiplierTest { 2, 40, },
		SmallMultiplierTest { 2, 999, },
		SmallMultiplierTest { 20, 40, },
		SmallMultiplierTest { 99, 99, },
	}

	for _, test := range SmallMultiplierTests {
		if HasPandigitalProduct(test.multiplier, test.multiplicand) {
			t.Errorf("%v x %v should not generate a pandigital product", test.multiplier, test.multiplicand)
		}
	}
}
