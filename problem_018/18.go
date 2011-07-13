package main

import (
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
			}

			rows[i + 1][j + 1] += item
		}
	}

	return Max(rows[len(rows) - 1])
}
