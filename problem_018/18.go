package main

type Element struct {
	value int
}

type ElementWithAdjacents struct {
	Element
	right, left *Element
}

func NewElement(value int, right, left *Element) *ElementWithAdjacents {
	element := new(ElementWithAdjacents)
	element.value = value
	element.right = right
	element.left = left
	return element
}
