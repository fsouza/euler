package main

import (
	"strconv"
	"strings"
)

type Element struct {
	value int
	right, left *Element
}

func (e *Element) String() string {
	return strconv.Itoa(e.value)
}

func NewElement(value int, left, right *Element) *Element {
	element := new(Element)
	element.value = value
	element.right = right
	element.left = left
	return element
}

func BuildFromString(input string) *Element {
	var root, left, right *Element

	rows := strings.Split(input, "\n", -1)
	for i := len(rows) - 1; i >= 0; i-- {
		row := rows[i]
		values := strings.Split(row, " ", -1)
		if len(values) == 2 {
			leftValue, _ := strconv.Atoi(values[0])
			rightValue, _ := strconv.Atoi(values[1])
			left = NewElement(leftValue, nil, nil)
			right = NewElement(rightValue, nil, nil)
		} else if len(values) == 1 {
			value, _ := strconv.Atoi(values[0])
			root = NewElement(value, left, right)
		}
	}

	return root
}
