package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const DayTwoTestInput = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

var Maximums = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

var NumberCapture = regexp.MustCompile(`(?P<number>\d+) (?P<color>red|green|blue)`)

func day2(part int, testing bool) {
	input := DayTwoTestInput
	if !testing {
		input = readInput(2, false)
	}

	sum := 0
Outer:
	for idx, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}

		id := idx + 1
		infoList := make([]map[string]string, 0)
		for _, match := range NumberCapture.FindAllStringSubmatch(line, -1) {
			info := make(map[string]string)
			for i, name := range NumberCapture.SubexpNames() {
				if i == 0 {
					continue
				}
				info[name] = match[i]
			}

			if part == 2 {
				infoList = append(infoList, info)
			} else {
				num, _ := strconv.Atoi(info["number"])
				maximum := Maximums[info["color"]]
				if num > maximum {
					continue Outer
				}
			}

		}
		if part == 2 {
			minimums := map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			}
			for _, info := range infoList {
				rolls, _ := strconv.Atoi(info["number"])
				if rolls > minimums[info["color"]] {
					minimums[info["color"]] = rolls
				}
			}
			prod := 1
			for _, num := range minimums {
				prod *= num
			}
			sum += prod
		} else {
			sum += id
		}
	}

	fmt.Printf("The sum is %d\n", sum)
}
