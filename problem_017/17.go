package main

var unitiesMap = map[int]string{
	1:"one",
	2:"two",
	3:"three",
	4:"four",
	5:"five",
	6:"six",
	7:"seven",
	8:"eight",
	9:"nine",
}

func WriteNumber(number int) string {
	return unitiesMap[number]
}
