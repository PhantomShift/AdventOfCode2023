package main

import (
	"fmt"
	"strconv"
	"strings"
)

const DayFourTestInput = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`

type set map[int]bool

func countIntersect(a set, b set) int {
	sum := 0
	for i := range a {
		if b[i] {
			sum++
		}
	}

	return sum
}

func day4BuildSet(numbers string) set {
	result := set{}
	for _, num := range strings.Split(numbers, " ") {
		num, err := strconv.Atoi(num)
		if err == nil {
			result[num] = true
		}
	}
	return result
}

func day4(part int, testing bool) {
	input := DayFourTestInput
	if !testing {
		input = readInput(4, false)
	}

	// don't feel like copy pasting all this code so if-else inside the loop lol
	instances := map[int]int{}
	points := 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		cardString, allNumbers, _ := strings.Cut(line, ":")
		cardStringSplit := strings.Split(cardString, " ")
		cardNumber, _ := strconv.Atoi(cardStringSplit[len(cardStringSplit)-1])
		instances[cardNumber]++
		winning, given, _ := strings.Cut(allNumbers, "|")
		winningSet := day4BuildSet(winning)
		givenSet := day4BuildSet(given)
		count := countIntersect(winningSet, givenSet)
		if count > 0 {
			if part == 1 {
				// apparently integer exponentiation isnt a thing in go???
				sum := 1
				for i := 0; i < count-1; i++ {
					sum *= 2
				}
				points += sum
			} else {
				copies := instances[cardNumber]
				for i := 1; i < count+1; i++ {
					// println(cardNumber + i)
					instances[cardNumber+i] += copies
				}
			}
		}
	}

	if part == 2 {
		for _, copies := range instances {
			points += copies
		}
	}

	fmt.Printf("The sum is %d\n", points)
}
