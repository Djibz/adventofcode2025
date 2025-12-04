package tools

import (
	"bufio"
	"fmt"
	"os"
)

func LinesChan(file *os.File) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()

	return ch
}

func GetMatrix(file *os.File) [][]byte {
	var matrix [][]byte

	for l := range LinesChan(file) {
		matrix = append(matrix, []byte(l))
	}

	return matrix
}

func PrintStringMatrix(matrix [][]byte) {
	for x := range len(matrix) {
		for y := range len(matrix[x]) {
			fmt.Print(string(matrix[x][y]))
		}
		fmt.Println()
	}
}
