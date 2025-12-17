package day9

import (
	"aoc/tools"
	"os"
	"regexp"
	"strconv"
)

type pouint struct {
	x int
	y int
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func size(a pouint, b pouint) int {
	return absInt(a.x-b.x+1) * absInt(a.y-b.y+1)
}

func First(file *os.File) int {
	pouints := []pouint{}
	r, _ := regexp.Compile(`(\d+),(\d+)`)
	for line := range tools.LinesChan(file) {
		res := r.FindStringSubmatch(line)
		x, _ := strconv.Atoi(res[1])
		y, _ := strconv.Atoi(res[2])
		pouints = append(pouints, pouint{x, y})
	}

	best := 0
	for _, p1 := range pouints {
		for _, p2 := range pouints {
			s := size(p1, p2)
			if s > best {
				best = s
			}
		}
	}

	return best
}

func Second(file *os.File) int {
	return 0
}
