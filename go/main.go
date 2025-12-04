package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
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
	file := openFile("../resources/1/input.txt")
	start := time.Now()
	fmt.Println("Day 1 step 1:", day1.First(file))
	fmt.Println("Time spent :", time.Since(start))

	file.Seek(0, 0)
	start = time.Now()
	fmt.Println("Day 1 step 2:", day1.Second(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Close()

	file = openFile("../resources/2/input.txt")
	start = time.Now()
	fmt.Println("Day 2 step 1:", day2.First(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Seek(0, 0)
	start = time.Now()
	fmt.Println("Day 2 step 2:", day2.Second(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Close()

	file = openFile("../resources/3/input.txt")
	start = time.Now()
	fmt.Println("Day 3 step 1:", day3.First(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Seek(0, 0)
	start = time.Now()
	fmt.Println("Day 3 step 2:", day3.Second(file))
	fmt.Println("Time spent :", time.Since(start))
	file.Close()
}
