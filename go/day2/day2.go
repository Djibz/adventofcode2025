package day2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func commaSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), ","); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if i := strings.Index(string(data), "\n"); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

func First(file *os.File) int {
	counter := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(commaSplit)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "-")
		first, _ := strconv.Atoi(values[0])
		second, _ := strconv.Atoi(values[1])

		for first <= second {
			str := strconv.Itoa(first)
			l := len(str)
			if l%2 == 0 {
				if str[0:l/2] == str[l/2:] {
					counter += first
				}
			}

			first++
		}
	}

	return counter
}

func FirstV2(file *os.File) int {
	counter := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(commaSplit)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "-")
		// first, _ := strconv.Atoi(values[0])
		second, _ := strconv.Atoi(values[1])

		half := values[0][:len(values[0])/2]
		half_int, _ := strconv.Atoi(half)
		fmt.Println(values, half)

		for {
			double, _ := strconv.Atoi(half + half)

			if double > second {
				break
			}

			fmt.Println(double)
			counter += double

			half_int++
			half = strconv.Itoa(half_int)
		}
	}

	return counter
}

func Second(file *os.File) int {
	counter := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(commaSplit)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "-")
		first, _ := strconv.Atoi(values[0])
		second, _ := strconv.Atoi(values[1])

		already := []int{}

		for first <= second {
			str := strconv.Itoa(first)

			l := 1
			for l <= len(str)/2 {
				r, _ := regexp.Compile(fmt.Sprintf(`^(%s)+$`, str[:l]))
				result := r.MatchString(str)

				if result && !slices.Contains(already, first) {
					// fmt.Println(first)
					already = append(already, first)
					counter += first
				}

				l++
			}

			first++
		}
	}

	return counter
}
