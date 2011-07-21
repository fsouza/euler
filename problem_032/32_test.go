package main

import (
	"testing"
)

type SmallMultiplierTest struct {
	multiplier, multiplicand int
}

func TestIfTheMultiplicandAndMultiplierHasLessThan5DigitsTogetherTheProductIsntPanDigital(t *testing.T) {
	SmallMultiplierTests := []SmallMultiplierTest {
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

func TestIsNotPandigitalSetIfThereIsARepeatedNumber(t *testing.T) {
	SmallMultiplierTests := []SmallMultiplierTest {
		SmallMultiplierTest { 222, 134, },
		SmallMultiplierTest { 123, 289, },
	}

	for _, test := range SmallMultiplierTests {
		if HasPandigitalProduct(test.multiplier, test.multiplicand) {
			t.Errorf("%v x %v should not generate a pandigital product", test.multiplier, test.multiplicand)
		}
	}
}

func TestIsNotPandigitalSetIfThereIsAZero(t *testing.T) {
	if HasPandigitalProduct(105, 234) {
		t.Errorf("%v x %v should not generate a pandigital product", 105, 234)
	}
}

