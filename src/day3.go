package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const DayThreeTestInput = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

type node struct {
	number   int
	numberId int
	char     rune
}

func buildGraph(input string) []map[int]node {
	graph := make([]map[int]node, 0, strings.Count(DayThreeTestInput, "\n"))
	currentNumberId := 1

	for _, line := range strings.Split(input, "\n") {
		mapLine := map[int]node{}
		currentSequence := make([]rune, 0, 3)
		indices := make([]int, 0, 3)

		updateNumbers := func() {
			if len(currentSequence) > 0 {
				number, _ := strconv.Atoi(string(currentSequence))
				for _, index := range indices {
					entry := mapLine[index]
					entry.number = number
					entry.numberId = currentNumberId
					mapLine[index] = entry
				}
				currentSequence = currentSequence[:0]
				indices = indices[:0]
				currentNumberId++
			}
		}

		for n, char := range line {
			if unicode.IsNumber(char) {
				currentSequence = append(currentSequence, char)
				indices = append(indices, n)
				mapLine[n] = node{
					number:   0,
					numberId: 0,
					char:     char,
				}
			} else {
				updateNumbers()
				if char != '.' {
					mapLine[n] = node{
						number:   -1,
						numberId: 0,
						char:     char,
					}
				}
			}
		}
		updateNumbers()
		graph = append(graph, mapLine)
	}

	return graph
}

func day3part1(graph []map[int]node) int {
	maxRows := len(graph)
	symbols := map[int]int{}

	for y, line := range graph {
		for x, node := range line {
			if node.number == -1 {
				checking := map[int]bool{
					y: true,
				}
				if y > 0 {
					checking[y-1] = true
				}
				if y < maxRows {
					checking[y+1] = true
				}
				for i := range checking {
					for _, k := range [3]int{x - 1, x, x + 1} {
						if other := graph[i][k]; other.number > 0 {
							if symbols[other.numberId] == 0 {
								symbols[other.numberId] = other.number
							}
						}
					}
				}
			}
		}
	}

	sum := 0
	for _, number := range symbols {
		sum += number
	}

	return sum
}

func day3part2(graph []map[int]node) int {
	maxRows := len(graph)
	sum := 0

	for y, line := range graph {
		for x, node := range line {
			if node.number == -1 {
				checking := map[int]bool{
					y: true,
				}
				if y > 0 {
					checking[y-1] = true
				}
				if y < maxRows {
					checking[y+1] = true
				}
				count := 0
				numbers := map[int]int{}
				for i := range checking {
					for _, k := range [3]int{x - 1, x, x + 1} {
						if other := graph[i][k]; other.number > 0 {
							if numbers[other.numberId] == 0 {
								count += 1
								numbers[other.numberId] = other.number
							}
						}
					}
				}
				if count == 2 {
					prod := 1
					for _, n := range numbers {
						prod *= n
					}
					sum += prod
				}
			}
		}
	}

	return sum
}

func day3(part int, testing bool) {
	input := DayThreeTestInput
	if !testing {
		input = readInput(3, false)
	}
	graph := buildGraph(input)

	sum := 0
	if part == 1 {
		sum = day3part1(graph)
	} else {
		sum = day3part2(graph)
	}

	fmt.Printf("The sum is %d\n", sum)
}
