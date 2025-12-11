package day7

import (
	"aoc/tools"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
)

func First(file *os.File) int {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()

	r, _ := regexp.Compile("S")
	a := r.FindStringIndex(text)

	beams := []int{a[0]}

	regex, _ := regexp.Compile(`\^`)
	splits := stepBeamsSplitsCount(beams, scanner, regex)

	return splits
}

func stepBeamsSplitsCount(beams []int, scanner *bufio.Scanner, r *regexp.Regexp) int {
	next_beams := []int{}
	if !(scanner.Scan()) {
		return 0
	}

	splitters := []int{}
	for _, i := range r.FindAllStringIndex(scanner.Text(), -1) {
		splitters = append(splitters, i[0])
	}

	splits := 0
	for _, b := range beams {
		if slices.Contains(splitters, b) {
			splits++
			if !slices.Contains(next_beams, b+1) {
				next_beams = append(next_beams, b+1)
			}
			if !slices.Contains(next_beams, b-1) {
				next_beams = append(next_beams, b-1)
			}
		} else {
			if !slices.Contains(next_beams, b) {
				next_beams = append(next_beams, b)
			}
		}
	}

	return stepBeamsSplitsCount(next_beams, scanner, r) + splits
}

func Second(file *os.File) int {
	lines := tools.GetMatrixString(file)

	r, _ := regexp.Compile("S")
	a := r.FindStringIndex(lines[0])

	regex, _ := regexp.Compile(`\^`)
	// TODO : préparer en avance les index des plitters au lieu de les rechercher à chaque fois

	return countPaths(a[0], &lines, 1, regex)
}

var cache = make(map[string]int)

func buildKey(beam int, line int) string {
	return fmt.Sprintf("%d-%d", beam, line)
}

func countPaths(beam int, lines *[]string, line int, r *regexp.Regexp) int {
	if line >= len(*lines) {
		return 1
	}

	key := buildKey(beam, line)
	if v, ok := cache[key]; ok {
		return v
	}

	splitters := []int{}
	for _, i := range r.FindAllStringIndex((*lines)[line], -1) {
		splitters = append(splitters, i[0])
	}

	if slices.Contains(splitters, beam) {
		subs := countPaths(beam-1, lines, line+1, r) + countPaths(beam+1, lines, line+1, r)
		cache[key] = subs
		return subs
	}

	subs := countPaths(beam, lines, line+1, r)
	cache[key] = subs
	return subs
}
