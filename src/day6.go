package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const DaySixTestInput = `Time:      7  15   30
Distance:  9  40  200
`

func readNumbers(s string) (func() bool, func() int) {
	i := 0
	current := ""
	next := func() bool {
		for i < len(s) && !unicode.IsDigit(rune(s[i])) {
			i++
		}
		if i >= len(s) {
			return false
		}
		size := 1
		for i+size < len(s) && unicode.IsDigit(rune(s[i+size])) {
			size++
		}
		current = s[i : i+size]
		i += size + 1
		return true
	}
	get := func() int {
		num, _ := strconv.Atoi(current)
		return num
	}
	return next, get
}

func collectNumbers(s string) []int {
	result := make([]int, 0)
	for next, get := readNumbers(s); next(); {
		result = append(result, get())
	}
	return result
}

func readCombinedNumber(s string) int {
	filtered := strings.Map(func(char rune) rune {
		if unicode.IsDigit(char) {
			return char
		}
		return -1
	}, s)
	n, _ := strconv.Atoi(filtered)
	return n
}

func day6(part int, testing bool) {
	input := DaySixTestInput
	if !testing {
		input = readInput(6, false)
	}

	timeLine, distanceLine, _ := strings.Cut(input, "\n")
	var times []int
	var distances []int
	if part == 1 {
		times = collectNumbers(timeLine)
		distances = collectNumbers(distanceLine)
	} else {
		times = []int{readCombinedNumber(timeLine)}
		distances = []int{readCombinedNumber(distanceLine)}
	}

	prod := 1
	for i, time := range times {
		distance := distances[i]
		for n := 1; n < time; n++ {
			if n*(time-n) > distance {
				prod *= (time - 2*n + 1)
				break
			}
		}
	}

	fmt.Printf("The product is %d\n", prod)
}
