package day4

import (
	"aoc/tools"
	"os"
)

func lazyGreaterThan(matrix *[][]byte, x int, y int, n int) bool {
	counter := 0
	for _, a := range []int{-1, 0, 1} {
		for _, b := range []int{-1, 0, 1} {
			if a == 0 && b == 0 {
				continue
			}
			x_seek := x + a
			y_seek := y + b

			if x_seek < 0 || y_seek < 0 || x_seek >= len(*matrix) || y_seek >= len((*matrix)[0]) {
				continue
			}

			if (*matrix)[x_seek][y_seek] == '@' {
				counter++
			}

			if counter > n {
				return true
			}
		}
	}

	return false
}

func First(file *os.File) int {
	matrix := tools.GetMatrix(file)

	counter := 0
	for x := range len(matrix) {
		for y := range len(matrix[0]) {
			if matrix[x][y] == '@' && !lazyGreaterThan(&matrix, x, y, 3) {
				counter++
			}
		}
	}

	return counter
}

func Second(file *os.File) int {
	matrix := tools.GetMatrix(file)

	counter := 0
	for {
		start := counter
		for x := range len(matrix) {
			for y := range len(matrix[0]) {
				if matrix[x][y] == '@' && !lazyGreaterThan(&matrix, x, y, 3) {
					counter++
					matrix[x][y] = '.'
				}
			}
		}
		if counter == start {
			break
		}
	}

	return counter
}
