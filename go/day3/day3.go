package day3

import (
	"aoc/tools"
	"os"
	"strconv"
)

func bestOfStr(str string) (index int, value string) {
	best := 0
	best_i := 0

	for i := 0; i < len(str); i++ {
		n, _ := strconv.Atoi(str[i : i+1])

		if n > best {
			best = n
			best_i = i
		}

		if best == 9 {
			break
		}
	}

	return best_i, strconv.Itoa(best)
}

func First(file *os.File) int {
	total := 0
	for line := range tools.LinesChan(file) {
		i1, v1 := bestOfStr(line[:len(line)-1])
		_, v2 := bestOfStr(line[i1+1:])

		value, _ := strconv.Atoi(v1 + v2)
		total += value
	}

	return total
}

func Second(file *os.File) int {
	total := 0
	for line := range tools.LinesChan(file) {

		begin := 0
		value := ""

		for r := range 12 {
			i, v := bestOfStr(line[begin : len(line)-(11-r)])

			begin = begin + i + 1
			value += v
		}

		int_val, _ := strconv.Atoi(value)
		total += int_val
	}

	return total
}
