package day6

import (
	"aoc/tools"
	"os"
	"regexp"
	"strconv"
)

func First(file *os.File) int {
	r, _ := regexp.Compile(`(\d+)`)

	numbers := [][]int{}
	operators := []byte{}
	for line := range tools.LinesChan(file) {
		if r.Match([]byte(line)) {
			number_line := []int{}
			for _, n_str := range r.FindAllStringSubmatch(line, -1) {
				n, _ := strconv.Atoi(n_str[0])
				number_line = append(number_line, n)
			}
			numbers = append(numbers, number_line)
		} else {
			r, _ := regexp.Compile(`(\+|\*)`)
			for _, operator := range r.FindAllStringSubmatch(line, -1) {
				operators = append(operators, []byte(operator[0])[0])
			}
		}
	}

	total := 0
	for i, operator := range operators {
		base := 0
		if operator == '*' {
			base = 1
		}
		for k := range len(numbers) {
			if operator == '+' {
				base += numbers[k][i]
			} else {
				base *= numbers[k][i]
			}
		}

		total += base
	}

	return total
}

func Second(file *os.File) int {
	return 0
}
