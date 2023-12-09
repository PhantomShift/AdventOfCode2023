package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const DayNineTestInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

func all_zeroes(values []int) bool {
	for _, v := range values {
		if v != 0 {
			return false
		}
	}
	return true
}

func traverse(values []int) [][]int {
	full := [][]int{values}
	for current := full[len(full)-1]; !all_zeroes(current); current = full[len(full)-1] {
		next := make([]int, len(current)-1)
		for i := 0; i < len(current)-1; i++ {
			next[i] = current[i+1] - current[i]
		}
		full = append(full, next)
	}

	return full
}

func extrapolate(values [][]int, backwards bool) int {
	values[len(values)-1] = append(values[len(values)-1], 0)
	if backwards {
		for _, v := range values {
			slices.Reverse(v)
		}
	}
	for i := len(values) - 1; i > 0; i-- {
		current := values[i]
		next := values[i-1]
		difference := current[len(current)-1]
		if backwards {
			difference = -difference
		}
		values[i-1] = append(next, difference+next[len(current)-1])
	}

	return values[0][len(values[0])-1]
}

func day9(part int, testing bool) {
	input := DayNineTestInput
	if !testing {
		input = readInput(9, false)
	}
	backwards := part == 2

	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		valStrings := strings.Split(line, " ")
		values := make([]int, len(valStrings))
		for i, s := range valStrings {
			values[i], _ = strconv.Atoi(s)
		}
		sum += extrapolate(traverse(values), backwards)
	}

	fmt.Printf("The sum is %d\n", sum)
}
