package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Row []int

func Max(items []int) (max int) {
	for _, item := range items {
		if item > max {
			max = item
		}
	}

	return
}

func BuildrowsSliceFromString(input string) []Row {
	rows := make([]Row, 0)
	stringRows := strings.Split(input, "\n")

	for _, stringRow := range stringRows {
		items := strings.Split(stringRow, " ")
		row := make([]int, len(items))
		for i, item := range items {
			row[i], _ = strconv.Atoi(item)
		}
		rows = append(rows, row)
	}

	return rows
}

func FindGreatestSum(rows []Row) int {
	for i, row := range rows[:len(rows) - 1] {
		for j, item := range row[:len(row)] {
			if j == 0 {
				rows[i + 1][j] += item
			} else if rows[i + 1][j] - row[j - 1] + item > rows[i + 1][j] {
				rows[i + 1][j] += item - row[j - 1]
			}

			rows[i + 1][j + 1] += item
		}
	}

	return Max(rows[len(rows) - 1])
}

func main() {
	input := `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`
	rows := BuildrowsSliceFromString(input)
	fmt.Println(FindGreatestSum(rows))
}
