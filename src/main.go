package main

import (
	"fmt"
	"os"
	"strconv"
)

func readInput(day int, testing bool) string {
	path := fmt.Sprintf("input/day%d", day)
	if testing {
		path += "test"
	}

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	return string(file)
}

func main() {
	args := os.Args
	if len(args) < 3 {
		println("Enter a day and part number")
		return
	}

	day, err := strconv.Atoi(args[1])
	if err != nil {
		println("Day should be a number")
		return
	} else if day < 1 || day > 25 {
		println("Day number should be between 1 and 25")
		return
	}
	part, err := strconv.Atoi(args[2])
	if err != nil || part != 1 && part != 2 {
		println("Part should be either 1 or 2")
		return
	}
	testing := false
	if len(args) > 3 {
		testing = args[3] == "test"
	}

	var dayFunc func(int, bool)
	switch day {
	case 1:
		// day1(part, testing)
		dayFunc = day1
	case 2:
		dayFunc = day2
	case 3:
		dayFunc = day3
	case 4:
		dayFunc = day4
	case 5:
		dayFunc = day5
	case 6:
		dayFunc = day6
	case 7:
		dayFunc = day7
	case 8:
		dayFunc = day8
	case 9:
		dayFunc = day9
	case 10:
		dayFunc = day10
	default:
		fmt.Printf("Day %d has not been implemented yet\n", day)
		return
	}
	dayFunc(part, testing)
}
