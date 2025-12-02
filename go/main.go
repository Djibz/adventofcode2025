package main

import (
	"aoc/day1"
	"aoc/day2"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("../resources/1/input.txt")
	check(err)
	fmt.Println("Day 1 step 1:", day1.First(file))
	file.Seek(0, 0)
	fmt.Println("Day 1 step 2:", day1.Second(file))
	file.Close()

	file, err = os.Open("../resources/2/input.txt")
	check(err)
	fmt.Println("Day 2 step 1:", day2.First(file))
	file.Seek(0, 0)
	fmt.Println("Day 2 step 2:", day2.Second(file))
}
