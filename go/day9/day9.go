package day9

import (
	"fmt"
	"os"
)

type pouint struct {
	x uint
	y uint
	z uint
}

func distance(a pouint, b pouint) uint {
	return uint(a.x-b.x) + uint(a.y-b.y) + uint(a.z-b.z)
}

func First(file *os.File) int {
	n := -3
	fmt.Println(uint(n))
	return 0
}

func Second(file *os.File) int {
	return 0
}
