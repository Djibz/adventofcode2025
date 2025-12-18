package day10

import (
	"aoc/tools"
	"fmt"
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
		fmt.Println(total)
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

func joltageToPattern(joltage []int) int {
	pattern := 0
	for i, v := range joltage {
		if v%2 != 0 {
			pattern |= 1 << i
		}
	}

	return pattern
}

type potential struct {
	state   int
	buttons []int
}

func eveningSolutions(initial int, goal int, buttons *[][]int) []int {
	currents := []potential{}

	solutions := []int{}

	for range len(*buttons) {
		next := []potential{}
		for _, p := range currents {
			for ib, b := range *buttons {
				new_state := clickButton(p.state, &b)
				new_potential := potential{new_state, append(p.buttons, ib)}
				if new_state == goal {
					solutions = append(solutions, new_potential.buttons...)
				} else {
					next = append(next, new_potential)
				}
			}
		}
		currents = next
	}

	return solutions
}

func shortestJoltage(goal []int, buttons *[][]int) int {
	if is0(goal) {
		return 0
	}

	shortest, clicked := horizontalShortestWithButtons(0, joltageToPattern(goal), buttons, []int{})
	fmt.Println(shortest)
	joltage := goal
	for _, button_index := range clicked {
		joltage = applyJoltageNegative(joltage, &(*buttons)[button_index])
	}

	return shortest + (2 * shortestJoltage(divideJoltage(joltage), buttons))
}
