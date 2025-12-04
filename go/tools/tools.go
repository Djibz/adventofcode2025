package tools

import (
	"bufio"
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
