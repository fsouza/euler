package main

import (
	"testing"
)

func TestTriangleElementHasValue(t *testing.T) {
	element := Element{68}
	if element.value != 68 {
		t.Errorf("Element should have a value!")
	}
}

