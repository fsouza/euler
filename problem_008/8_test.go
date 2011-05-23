package main

import (
	"testing"
)

type FindProductTest struct {
	in string
	out int
}

var FindProductTests = []FindProductTest{
	FindProductTest{"1234567890", 15120},
	FindProductTest{"2345678999993928932", 59049},
	FindProductTest{"123499999", 59049},
}

func TestFindProduct(t *testing.T) {
	for _, v := range FindProductTests {
		got := FindProduct(v.in)
		if v.out != got {
			t.Errorf("I was expecting the greatest product to be %d, but it was %d :(", v.out, got)
		}
	}
}
