package main

import (
	"testing"
)

func AssertAreSliceEqual(expected, got []int, t *testing.T) {
	if len(expected) != len(got) {
		t.Errorf("Assertion error.\nExpected: %v\nGot:%v", expected, got)
	} else {
		for i := 0; i < len(expected); i++ {
			if expected[i] != got[i] {
				t.Errorf("Assertion error.\nExpected: %v\nGot:%v", expected, got)
			}
		}
	}
}

func TestBuildSliceOfPrimes(t *testing.T) {
	expected := []int{ 2, 3, 5, 7, }
	got := BuildSliceOfPrimes(10)
	AssertAreSliceEqual(expected, got, t)
}

