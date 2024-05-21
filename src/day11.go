package main

import (
	"fmt"
	"slices"
	"strings"
)

const DayElevenTestInput = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`

func expandSpace(space string) string {
	result := strings.Split(space, "\n")
	emptyLine := strings.Repeat(".", len(result[0]))
	for i := 0; i < len(result); i++ {
		if len(result[i]) == 0 {
			continue
		}
		if strings.Count(result[i], "#") == 0 {
			result = slices.Insert(result, i+1, emptyLine)
			i++
		}
	}
	for i := 0; i < len(result[0]); i++ {
		isEmpty := true
		for _, line := range result {
			if len(line) == 0 {
				continue
			}
			if line[i] == '#' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			for k, line := range result {
				if len(line) == 0 {
					continue
				}
				result[k] = line[:i] + "." + line[i:]
			}
			i++
		}
	}

	return strings.Join(result, "\n")
}

func getPositions(space string) []position {
	positions := []position{}
	for y, line := range strings.Split(space, "\n") {
		for x, char := range line {
			if char == '#' {
				positions = append(positions, position{x: x, y: y})
			}
		}
	}

	return positions
}

func sumDistances(positions []position) int {
	// too lazy to do this properly so just checkedPairs[min][max]
	checkedPairs := map[int]map[int]bool{}

	sum := 0
	for a, posA := range positions {
		for b, posB := range positions {
			lesser, greater := min(a, b), max(a, b)
			if checkedPairs[lesser] == nil {
				checkedPairs[lesser] = map[int]bool{}
			}
			if a == b || checkedPairs[lesser][greater] {
				continue
			}
			checkedPairs[lesser][greater] = true
			distance := max(posA.x, posB.x) + max(posA.y, posB.y) - min(posA.x, posB.x) - min(posA.y, posB.y)
			sum += distance
		}
	}

	return sum
}

func expandAndSumSpace(space string, expansionRate int) int {
	positions := getPositions(space)
	checkedPairs := map[int]map[int]bool{}
	emptyColumns := map[int]bool{}
	emptyRows := map[int]bool{}
	splitted := strings.Split(space, "\n")
	for r, line := range splitted {
		if strings.Count(line, "#") == 0 {
			emptyRows[r] = true
		}
	}
outer:
	for c := 0; c < len(splitted[0]); c++ {
		for _, line := range splitted {
			if len(line) == 0 {
				continue
			}
			if line[c] == '#' {
				continue outer
			}
		}
		emptyColumns[c] = true
	}

	sum := 0
	for a, posA := range positions {
		for b, posB := range positions {
			lesser, greater := min(a, b), max(a, b)
			if checkedPairs[lesser] == nil {
				checkedPairs[lesser] = map[int]bool{}
			}
			if a == b || checkedPairs[lesser][greater] {
				continue
			}
			checkedPairs[lesser][greater] = true
			distance := 0
			for x := min(posA.x, posB.x); x < max(posA.x, posB.x); x++ {
				if emptyColumns[x] {
					distance += expansionRate
				} else {
					distance++
				}
			}
			for y := min(posA.y, posB.y); y < max(posA.y, posB.y); y++ {
				if emptyRows[y] {
					distance += expansionRate
				} else {
					distance++
				}
			}
			sum += distance
		}
	}

	return sum
}

func day11(part int, testing bool) {
	input := DayElevenTestInput
	if !testing {
		input = readInput(11, false)
	}
	if part == 1 {
		fmt.Printf("The sum is %d\n", sumDistances(getPositions(expandSpace(input))))
	} else {
		// although the work for part 2 applies to part 1 with a rate of 2,
		// it seems to be significantly slower for the real input on my machine
		rate := 100
		if !testing {
			rate = 1000000
		}
		fmt.Printf("The sum is %d\n", expandAndSumSpace(input, rate))
	}
}
