package main

import (
	"testing"
)

func AssertEqual(expected, got int, t *testing.T) {
	if expected != got {
		t.Errorf("Assertion error.\nExpected: %v.\nGot: %v.", expected, got)
	}
}
