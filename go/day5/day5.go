package day5

import (
	"aoc/tools"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func indexesFromTo(from int, to int, ingredients []int) (s int, e int) {
	start := -1
	end := -1
	for k, v := range ingredients {
		if start == -1 && v >= from {
			start = k
		}

		if end == -1 && start != -1 && v > to {
			end = k
			break
		}
	}

	return start, end
}

func First(file *os.File) int {
	ranges := [][]int{}
	ingredients := []int{}

	first := true
	for line := range tools.LinesChan(file) {
		if line == "" {
			first = false
			continue
		}

		if first {
			values := strings.Split(line, "-")
			from, _ := strconv.Atoi(values[0])
			to, _ := strconv.Atoi(values[1])

			ranges = append(ranges, []int{from, to})
		} else {
			number, _ := strconv.Atoi(line)
			ingredients = append(ingredients, number)
		}
	}

	slices.Sort(ingredients)
	fmt.Println(ranges)
	fmt.Println(ingredients)

	total := 0
	for _, r := range ranges {
		s, e := indexesFromTo(r[0], r[1], ingredients)

		ingredients = slices.Delete(ingredients, s, e)
		total += e - s
	}

	return total
}

func Second(file *os.File) int {
	ranges := [][]int{}

	for line := range tools.LinesChan(file) {
		if line == "" {
			break
		}

		values := strings.Split(line, "-")
		from, _ := strconv.Atoi(values[0])
		to, _ := strconv.Atoi(values[1])

		ranges = append(ranges, []int{from, to})
	}

	return 0
}

func addDomain(domains [][]int, singleton []int) [][]int {

}
