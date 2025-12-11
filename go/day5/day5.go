package day5

import (
	"aoc/tools"
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

	total := 0
	for _, r := range ranges {
		s, e := indexesFromTo(r[0], r[1], ingredients)

		ingredients = slices.Delete(ingredients, s, e)
		total += e - s
	}

	return total
}

func Second(file *os.File) int {
	ranges := []int{}

	for line := range tools.LinesChan(file) {
		if line == "" {
			break
		}

		values := strings.Split(line, "-")
		from, _ := strconv.Atoi(values[0])
		to, _ := strconv.Atoi(values[1])

		ranges = addDomain(ranges, from, to)
	}

	total := 0
	for i := range ranges {
		if i%2 != 0 {
			total += ranges[i] - ranges[i-1] + 1
		}
	}

	return total
}

func addDomain(domains []int, from int, to int) []int {
	i1 := placeToBe(domains, from)
	i2 := placeToBe(domains, to)
	l := len(domains)

	i2_even := i2%2 == 0
	i2_equal := i2 < l && to == domains[i2]

	if i1 < l && from == domains[i1] {
		if i1%2 == 0 {
			// in valid -> nothing to do
		} else {
			// use previous start
			i1--
		}
	} else {
		if i1%2 == 0 {
			// in invalid -> create new
			domains = slices.Insert(domains, i1, from)
			// push everything
			i2++
		} else {
			// in valid -> use previous
			i1--
		}
	}

	if i2_equal {
		if i2_even {
			// start of range -> use end
			i2++
		} else {
			// end of range -> use it
		}
	} else {
		if i2_even {
			// insert new
			domains = slices.Insert(domains, i2, to)
		} else {
			// in valid-> use it
		}
	}

	return slices.Delete(domains, i1+1, i2)
}

func placeToBe(list []int, value int) int {
	for i, v := range list {
		if v >= value {
			return i
		}
	}

	return len(list)
}
