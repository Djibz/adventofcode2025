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
	matrix := tools.GetMatrix(file)
	op_i := len(matrix) - 1

	total := 0

	current_index := 0
	for {
		operator := matrix[op_i][current_index]
		length := segmentLength(&matrix[op_i], current_index)

		result := 0
		if operator == '*' {
			result = 1
		}

		for offset := range length {
			str_number := ""
			for l_i := range len(matrix) - 1 {
				char := matrix[l_i][current_index+offset]
				if char != ' ' {
					str_number += string(char)
				}
			}
			number, _ := strconv.Atoi(str_number)
			if operator == '+' {
				result += number
			} else {
				result *= number
			}
		}

		total += result
		current_index += length + 1
		if current_index >= len(matrix[op_i]) {
			break
		}
	}

	return total
}

func segmentLength(matrix *[]byte, current int) int {
	length := 1
	for current+length < len(*matrix) &&
		(*matrix)[current+length] != '*' &&
		(*matrix)[current+length] != '+' {
		length++
	}

	if current+length == len(*matrix) {
		return length
	}

	return length - 1
}
