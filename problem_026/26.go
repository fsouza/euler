package main

import (
	"fmt"
	"strings"
)

func ExtractDecimal(number float64) string {
	out := fmt.Sprintf("%v", number)
	decimal := strings.Split(out, ".")[1]
	return decimal
}
