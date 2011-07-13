package main

type Row []int

func Max(items []int) (max int) {
	for _, item := range items {
		if item > max {
			max = item
		}
	}

	return
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
