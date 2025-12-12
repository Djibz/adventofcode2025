package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"aoc/day4"
	"aoc/day5"
	"aoc/day6"
	"aoc/day7"
	"aoc/day8"
	"fmt"
	"os"
	"time"
)

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return file
}

func main() {
	hello := time.Now()

	file := openFile("../resources/1/input.txt")
	start := time.Now()
	fmt.Println("Day 1 step 1:", day1.First(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Seek(0, 0)
	start = time.Now()
	fmt.Println("Day 1 step 2:", day1.Second(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Close()
	fmt.Println()

	file = openFile("../resources/2/example.txt")
	start = time.Now()
	fmt.Println("Day 2 step 1:", day2.First(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Seek(0, 0)
	start = time.Now()
	fmt.Println("Day 2 step 2:", day2.Second(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Close()
	fmt.Println()

	file = openFile("../resources/3/input.txt")
	start = time.Now()
	fmt.Println("Day 3 step 1:", day3.First(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Seek(0, 0)
	start = time.Now()
	fmt.Println("Day 3 step 2:", day3.Second(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Close()
	fmt.Println()

	file = openFile("../resources/4/input.txt")
	start = time.Now()
	fmt.Println("Day 4 step 1:", day4.First(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Seek(0, 0)
	start = time.Now()
	fmt.Println("Day 4 step 2:", day4.Second(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Close()
	fmt.Println()

	file = openFile("../resources/5/input.txt")
	start = time.Now()
	fmt.Println("Day 5 step 1:", day5.First(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Seek(0, 0)
	start = time.Now()
	fmt.Println("Day 5 step 2:", day5.Second(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Close()
	fmt.Println()

	file = openFile("../resources/6/input.txt")
	start = time.Now()
	fmt.Println("Day 6 step 1:", day6.First(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Seek(0, 0)
	start = time.Now()
	fmt.Println("Day 6 step 2:", day6.Second(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Close()
	fmt.Println()

	file = openFile("../resources/7/input.txt")
	start = time.Now()
	fmt.Println("Day 7 step 1:", day7.First(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Seek(0, 0)
	start = time.Now()
	fmt.Println("Day 7 step 2:", day7.Second(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Close()
	fmt.Println()

	file = openFile("../resources/8/input.txt")
	start = time.Now()
	fmt.Println("Day 8 step 1:", day8.First(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Seek(0, 0)
	start = time.Now()
	fmt.Println("Day 8 step 2:", day8.Second(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Close()
	fmt.Println()

	fmt.Println("Total time spent :", time.Since(hello))
}
