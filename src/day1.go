package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var NumberMapping = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func getDigitIncludingWords(input string, backwards bool) string {
	start := 0
	finish := len(input) - 1
	step := 1
	if backwards {
		start, finish = finish, start
		step = -1
	}

	for i := start; i != finish+step; i += step {
		if unicode.IsDigit(rune(input[i])) {
			return input[i : i+1]
		}
		checking := input[min(start, i):max(start+1, i+1)]
		for word, number := range NumberMapping {
			if strings.Contains(checking, word) {
				return number
			}
		}
	}

	return ""
}

func day1(part int, testing bool) {
	input := readInput(1, testing)

	getDigit := getDigitIncludingWords
	if part == 1 {
		input = strings.Map(func(char rune) rune {
			if unicode.IsDigit(char) || char == '\n' {
				return char
			}
			return -1
		}, input)
		getDigit = func(line string, backwards bool) string {
			if backwards {
				return line[len(line)-1:]
			}
			return line[0:1]
		}
	}

	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) < 1 {
			continue
		}
		left := getDigit(line, false)
		right := getDigit(line, true)
		num, err := strconv.Atoi(left + right)
		if err != nil {
			println("Error reading file:", err.Error())
			return
		}
		sum += num
	}

	fmt.Printf("The sum is %d\n", sum)
}
