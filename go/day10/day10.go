package day10

import (
	"aoc/tools"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func getButtons(line string) [][]int {
	buttons := [][]int{}
	pattern := regexp.MustCompile(`\([0-9,]+\)`)
	digits := regexp.MustCompile(`\d+`)
	for _, strButton := range pattern.FindAllString(line, -1) {
		button := []int{}
		for _, nStr := range digits.FindAllString(strButton, -1) {
			n, _ := strconv.Atoi(nStr)
			button = append(button, n)
		}
		buttons = append(buttons, button)
	}

	return buttons
}

func indicatorToInt(indicator string) int {
	number := 0
	for i, c := range indicator {
		if c == '#' {
			number |= 1 << i
		}
	}

	return number
}

func clickButton(state int, button *[]int) int {
	for _, n := range *button {
		state ^= 1 << n
	}

	return state
}

func First(file *os.File) int {
	total := 0
	re_indicator := regexp.MustCompile(`\[([\.#]+)]`)
	for line := range tools.LinesChan(file) {
		indicator := re_indicator.FindStringSubmatch(line)[1]
		buttons := getButtons(line)
		i_indicator := indicatorToInt(indicator)
		total += horizontalShortest(0, i_indicator, &buttons)
	}

	return total
}

func horizontalShortest(initial int, goal int, buttons *[][]int) int {
	explored := []int{}
	currents := []int{initial}
	depth := 0
	for {
		depth++
		next := []int{}

		for _, state := range currents {
			for _, button := range *buttons {
				newState := clickButton(state, &button)
				if newState == goal {
					return depth
				}
				if !slices.Contains(explored, newState) {
					explored = append(explored, newState)
					next = append(next, newState)
				}
			}
		}

		currents = next
	}
}

func Second(file *os.File) int {
	return 0
}
