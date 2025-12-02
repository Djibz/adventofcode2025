package day1

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func First(file *os.File) int {
	position := 50
	counter := 0

	reader := bufio.NewReader(file)
	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		r, _ := regexp.Compile(`(L|R)(\d+)`)
		values := r.FindStringSubmatch(string(bytes))

		forward := values[1] == "R"
		number, _ := strconv.Atoi(values[2])

		if forward {
			position = (position + number) % 100
		} else {
			position = (100 + position - number) % 100
		}

		if position == 0 {
			counter++
		}
	}

	return counter
}

func Second(file *os.File) int {
	position := 50
	counter := 0

	reader := bufio.NewReader(file)
	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		r, _ := regexp.Compile(`(L|R)(\d+)`)
		values := r.FindStringSubmatch(string(bytes))

		forward := values[1] == "R"
		number, _ := strconv.Atoi(values[2])

		div := number / 100
		counter += div
		number %= 100

		if forward {
			position += number
		} else {
			position -= number
		}

		if position < 0 {
			if position+number != 0 {
				counter++
			}
			position += 100
		}

		if position == 0 {
			counter++
		}

		if position >= 100 {
			counter++
			position -= 100
		}
	}

	return counter
}
