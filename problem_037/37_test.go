package main

import (
	"testing"
)

func TestIsTruncatable(t *testing.T) {
	primes := BuildPrimesList(3798)
	expected := true
	got := IsTruncatable(3797, primes)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

func TestTruncate(t *testing.T) {
	expected := 123
	got := Truncate(1234, "right")
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}


func TestTruncateRight(t *testing.T) {
	expected := 123
	got := TruncateRight(1234)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

func TestTruncateLeft(t *testing.T) {
	expected := 234
	got := TruncateLeft(1234)
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}

