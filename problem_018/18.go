package main

import (
	"strconv"
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
