package main

import (
	"fmt"
	"strings"
)

type mapNode struct {
	left  string
	right string
}

type mapNetwork map[string]mapNode

const DayEightTestInput = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`

const DayEightPartTwoTestInput = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`

// gcd implementation based on Wikipedia description
// https://en.wikipedia.org/wiki/Euclidean_algorithm
func gcd(a int, b int) int {
	if a == b {
		return a
	}

	big := max(a, b)
	small := min(a, b)
	remainder := big % small
	if remainder == 0 {
		return small
	}

	return gcd(small, remainder)
}

func day8(part int, testing bool) {
	input := DayEightTestInput
	if !testing {
		input = readInput(8, false)
	} else if part == 2 {
		input = DayEightPartTwoTestInput
	}

	instructions, list, _ := strings.Cut(input, "\n\n")

	solver := day8part1
	if part == 2 {
		solver = day8part2
	}
	steps := solver(instructions, list)

	fmt.Printf("It took %d steps\n", steps)
}

func day8part1(instructions string, list string) int {
	network := mapNetwork{}
	for _, line := range strings.Split(list, "\n") {
		if len(line) == 0 {
			continue
		}

		label, directions, _ := strings.Cut(line, " = (")
		left := directions[0:3]
		right := directions[5:8]
		network[label] = mapNode{left: left, right: right}
	}

	current := "AAA"
	steps := 0
	for current != "ZZZ" {
		for _, direction := range instructions {
			next := network[current].left
			if direction == 'R' {
				next = network[current].right
			}
			steps++
			current = next
			if current == "ZZZ" {
				break
			}
		}
	}

	return steps
}

func day8part2(instructions string, list string) int {
	network := mapNetwork{}
	starts := make([]string, 0)
	ends := map[string]bool{}
	for _, line := range strings.Split(list, "\n") {
		if len(line) == 0 {
			continue
		}

		label, directions, _ := strings.Cut(line, " = (")
		if last := rune(label[2]); last == 'A' {
			starts = append(starts, label)
		} else if last == 'Z' {
			ends[label] = true
		}
		left := directions[0:3]
		right := directions[5:8]
		network[label] = mapNode{left: left, right: right}
	}

	mins := make([]int, len(starts))
	for i, current := range starts {
		steps := 0
		for rune(current[2]) != 'Z' {
			for _, direction := range instructions {
				next := network[current].left
				if direction == 'R' {
					next = network[current].right
				}
				steps++
				current = next
				if current == "ZZZ" {
					break
				}
			}
		}
		mins[i] = steps
	}

	lcm := mins[0]
	for i := 1; i < len(mins); i++ {
		n := mins[i]
		lcm = lcm * n / gcd(lcm, n)
	}

	return lcm
}
