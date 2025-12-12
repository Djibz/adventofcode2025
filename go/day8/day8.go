package day8

import (
	"aoc/tools"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type pouint struct {
	x int
	y int
	z int
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func distance(a pouint, b pouint) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2) + math.Pow(float64(a.z-b.z), 2))
}

func createMatrix(pouints []pouint) [][]float64 {
	matrix := [][]float64{}
	for i, a := range pouints {
		matrix = append(matrix, []float64{})
		for j, b := range pouints {
			if j >= i {
				break
			}

			matrix[i] = append(matrix[i], distance(a, b))
		}
	}

	return matrix
}

func getClosests(matrix *[][]float64) (int, int) {
	bi := 0
	bj := 0
	best := 0.0

	for i, line := range *matrix {
		for j, distance := range line {
			if distance != 0 && (best == 0 || distance < best) {
				bi = i
				bj = j
				best = distance
			}
		}
	}

	return bi, bj
}

func First(file *os.File) int {
	pouints := []pouint{}
	groups := []int{}

	r, _ := regexp.Compile(`(\d+),(\d+),(\d+)`)
	for line := range tools.LinesChan(file) {
		res := r.FindStringSubmatch(line)
		x, _ := strconv.Atoi(res[1])
		y, _ := strconv.Atoi(res[2])
		z, _ := strconv.Atoi(res[3])

		pouints = append(pouints, pouint{x, y, z})
		groups = append(groups, 0)
	}
	distance_matrix := createMatrix(pouints)

	n_loops := 10
	if len(pouints) > 100 {
		n_loops = 1000
	}
	for range n_loops {
		i, j := getClosests(&distance_matrix)
		distance_matrix[i][j] = 0

		if groups[i] == groups[j] && groups[i] != 0 {
			continue
		}
		group := max(groups[i], groups[j])
		other := min(groups[i], groups[j])
		if group == 0 {
			group = slices.Max(groups) + 1
			groups[i] = group
			groups[j] = group
		} else if other == 0 {
			groups[i] = group
			groups[j] = group
		} else {
			for id, v := range groups {
				if v == other {
					groups[id] = group
				}
			}
		}
	}

	count := make(map[int]int)
	for _, v := range groups {
		count[v] += 1
	}

	counts := []int{}
	for k, c := range count {
		if k == 0 {
			continue
		}
		counts = append(counts, c)
	}
	slices.Sort(counts)

	total := 1
	for _, n := range counts[len(counts)-3:] {
		total *= n
	}

	return total
}

func createMap(groups []int) map[int]int {
	count := make(map[int]int)
	for _, v := range groups {
		count[v] += 1
	}

	return count
}

func Second(file *os.File) int {
	pouints := []pouint{}
	groups := []int{}

	r, _ := regexp.Compile(`(\d+),(\d+),(\d+)`)
	for line := range tools.LinesChan(file) {
		res := r.FindStringSubmatch(line)
		x, _ := strconv.Atoi(res[1])
		y, _ := strconv.Atoi(res[2])
		z, _ := strconv.Atoi(res[3])

		pouints = append(pouints, pouint{x, y, z})
		groups = append(groups, 0)
	}
	distance_matrix := createMatrix(pouints)

	for {
		i, j := getClosests(&distance_matrix)
		distance_matrix[i][j] = 0

		if groups[i] == groups[j] && groups[i] != 0 {
			continue
		}
		group := max(groups[i], groups[j])
		other := min(groups[i], groups[j])
		if group == 0 {
			group = slices.Max(groups) + 1
			groups[i] = group
			groups[j] = group
		} else if other == 0 {
			groups[i] = group
			groups[j] = group
		} else {
			for id, v := range groups {
				if v == other {
					groups[id] = group
				}
			}
		}

		if len(createMap(groups)) == 1 {
			return pouints[i].x * pouints[j].x
		}
	}
}
