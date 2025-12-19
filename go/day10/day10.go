package day10

import (
	"aoc/tools"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
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

func intToIndicator(number int) string {
	if number == 0 {
		return ""
	}

	char := ""
	if number%2 == 0 {
		char = "."
	} else {
		char = "#"
		number--
	}

	return intToIndicator(number/2) + char
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
	total := 0
	re_joltage := regexp.MustCompile(`\{(.+)\}`)
	for line := range tools.LinesChan(file) {
		joltage := re_joltage.FindStringSubmatch(line)[1]

		buttons := getButtons(line)
		target_joltage := joltageToInt(joltage)

		total += shortestJoltage(target_joltage, &buttons)
	}

	return total
}

func horizontalShortestWithButtons(initial int, goal int, buttons *[][]int, clicked []int) (int, []int) {
	explored := []int{}
	currents := []int{initial}
	depth := 0

	for {
		depth++
		next := []int{}

		for _, state := range currents {
			for i, button := range *buttons {
				newState := clickButton(state, &button)
				if newState == goal {
					return depth, append(clicked, i)
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

func joltageToInt(j string) []int {
	joltage := []int{}
	for s := range strings.SplitSeq(j, ",") {
		n, _ := strconv.Atoi(s)
		joltage = append(joltage, n)
	}

	return joltage
}

func applyJoltageNegative(state []int, button *[]int) []int {
	new := make([]int, len(state))
	copy(new, state)
	for _, n := range *button {
		new[n]--
	}

	return new
}

func divideJoltage(state []int) []int {
	new := []int{}
	copy(new, state)
	for _, n := range state {
		new = append(new, n/2)
	}

	return new
}

func is0(a []int) bool {
	for _, n := range a {
		if n != 0 {
			return false
		}
	}

	return true
}

func isNegative(a []int) bool {
	for _, n := range a {
		if n < 0 {
			return true
		}
	}

	return false
}

func joltageToPattern(joltage []int) int {
	pattern := 0
	for i, v := range joltage {
		if v%2 != 0 {
			pattern |= 1 << i
		}
	}

	return pattern
}

func combinaisons(size int) [][]bool {
	if size == 1 {
		return [][]bool{{false}, {true}}
	}

	combs := [][]bool{}
	for _, l := range combinaisons(size - 1) {
		n := []bool{false}
		n = append(n, l...)
		combs = append(combs, n)

		n2 := []bool{true}
		n2 = append(n2, l...)
		combs = append(combs, n2)
	}
	return combs
}

func applyCombinaison(combinaison []bool, state int, buttons *[][]int) int {
	for i, v := range combinaison {
		if v {
			for _, id := range (*buttons)[i] {
				state ^= 1 << id
			}
		}
	}

	return state
}

func eveningSolutions(initial int, buttons *[][]int) [][]int {
	solutions := [][]int{}

	for _, combinaison := range combinaisons(len(*buttons)) {
		if applyCombinaison(combinaison, initial, buttons) == 0 {
			solution := []int{}
			for i, v := range combinaison {
				if v {
					solution = append(solution, i)
				}
			}
			solutions = append(solutions, solution)
		}
	}

	return solutions
}

func shortestJoltage(goal []int, buttons *[][]int) int {
	if is0(goal) {
		return 0
	}

	sizes := []int{}

	for _, clicked := range eveningSolutions(joltageToPattern(goal), buttons) {
		joltage := goal
		for _, button_index := range clicked {
			joltage = applyJoltageNegative(joltage, &(*buttons)[button_index])
		}
		if isNegative(joltage) {
			continue
		}

		sub := shortestJoltage(divideJoltage(joltage), buttons)
		if sub != -1 {
			sizes = append(sizes, len(clicked)+(2*sub))
		}
	}

	if len(sizes) == 0 {
		return -1
	}

	return slices.Min(sizes)
}
