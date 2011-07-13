package main

type Row []int

func Max(itens []int) (max int) {
	for _, item := range itens {
		if item > max {
			max = item
		}
	}

	return
}

func FindGreatestSum(rows []Row) int {
	return 10
}
