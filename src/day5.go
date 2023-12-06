package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type rangeMapping struct {
	input  int
	output int
	size   int
}

type numRange struct {
	start int
	size  int
}

// Returns -1 if `number` is not in the range of `input`,
// otherwise returns the corresponding output
func (self *rangeMapping) getMapped(number int) int {
	if number < self.input || number >= self.input+self.size {
		return -1
	}

	return self.output + number - self.input
}

func (self *rangeMapping) getMappedRange(r numRange) []numRange {
	println("input: ", r.start, r.size)
	println("mapping:", self.input, self.output, self.size)
	result := make([]numRange, 0)
	if r.start < self.input {
		result = append(result, numRange{
			start: r.start,
			size:  min(r.size, self.input-r.start),
		})
	}
	if r.start+r.size >= self.input && r.start < self.input {
		result = append(result, numRange{
			start: self.output,
			size:  min(self.size, r.start+r.size-self.input),
		})
	}
	if r.start < self.input+self.size && r.start+r.size > self.input+self.size {
		result = append(result, numRange{
			start: r.start + r.size,
			size:  r.start + r.size - self.input - self.size,
		})
	}

	if len(result) == 0 {
		result = []numRange{r}
	}

	for _, r := range result {
		println(r.start, r.size)
	}

	return result
}

func atoiUnchecked(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}

func day5(part int, testing bool) {
	input := readInput(5, testing)

	seedsString, input, _ := strings.Cut(input, "\n\n")
	seedsStrings := strings.Split(seedsString, " ")
	mapInfo := strings.Split(input, "\n\n")
	mappings := make([][]rangeMapping, 0)
	for _, info := range mapInfo {
		info := strings.Split(info, "\n")
		current := make([]rangeMapping, 0, len(info))
		for i := 1; i < len(info); i++ {
			line := info[i]
			if len(line) == 0 {
				continue
			}
			split := strings.Split(line, " ")
			current = append(current, rangeMapping{
				input:  atoiUnchecked(split[1]),
				output: atoiUnchecked(split[0]),
				size:   atoiUnchecked(split[2]),
			})
		}
		mappings = append(mappings, current)
	}

	getLocation := func(seed int) int {
		current := seed
		for _, layer := range mappings {
			for _, mapping := range layer {
				mapped := mapping.getMapped(current)
				if mapped != -1 {
					current = mapped
					break
				}
			}
		}
		return current
	}

	minimum := math.MaxInt
	if part == 1 {
		for _, seed := range seedsStrings {
			seed, err := strconv.Atoi(seed)
			if err == nil {
				location := getLocation(seed)
				if location < minimum {
					minimum = location
				}
			}
		}
		// unfinished optimized solution evaluated via ranges
		// } else {
		// 	ranges := make([]numRange, 0, (len(seedsStrings)-1)/2)
		// 	for i := 1; i < len(seedsStrings)-1; i += 2 {
		// 		start, size := atoiUnchecked(seedsStrings[i]), atoiUnchecked(seedsStrings[i+1])
		// 		ranges = append(ranges, numRange{start: start, size: size})
		// 	}
		// 	for _, layer := range mappings {
		// 		nextRanges := make([]numRange, 0, len(ranges))
		// 		for _, mapping := range layer {
		// 			for len(ranges) > 0 {
		// 				r := ranges[0]
		// 				ranges = ranges[1:]
		// 				for _, new := range mapping.getMappedRange(r) {
		// 					nextRanges = append(nextRanges, new)
		// 				}
		// 			}
		// 		}
		// 		ranges = nextRanges
		// 	}
		// 	for _, r := range ranges {
		// 		// println(r.start, r.size)
		// 		if r.start < minimum {
		// 			minimum = r.start
		// 		}
		// 	}
		// }
	} else {
		for i := 1; i < len(seedsStrings)-1; i += 2 {
			start, size := atoiUnchecked(seedsStrings[i]), atoiUnchecked(seedsStrings[i+1])
			for add := 0; add < size; add++ {
				location := getLocation(start + add)
				if location < minimum {
					minimum = location
				}
			}
		}
	}

	fmt.Printf("The lowest location number is %d\n", minimum)
}
