package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

const DayTenTestInput = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`

const DayTenPartTwoTestInput = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
`

func clamp[T cmp.Ordered](x, low, high T) T {
	return max(low, min(x, high))
}

type position struct{ x, y int }

var TileMapping = map[rune][]position{
	'|': {
		position{x: 0, y: -1},
		position{x: 0, y: 1},
	},
	'-': {
		position{x: -1, y: 0},
		position{x: 1, y: 0},
	},
	'L': {
		position{x: 0, y: -1},
		position{x: 1, y: 0},
	},
	'J': {
		position{x: 0, y: -1},
		position{x: -1, y: 0},
	},
	'7': {
		position{x: 0, y: 1},
		position{x: -1, y: 0},
	},
	'F': {
		position{x: 0, y: 1},
		position{x: 1, y: 0},
	},
	'S': {
		position{x: 0, y: -1},
		position{x: 0, y: 1},
		position{x: -1, y: 0},
		position{x: 1, y: 0},
	},
}

type pipe struct {
	id        int
	tile      rune
	neighbors []int
}

type pipeField struct {
	start int
	x, y  int
	pipes map[int]pipe
}

func getId(maxX, x, y int) int {
	return maxX*y + x
}

func createPipeField(input string) pipeField {
	pipes := map[int]pipe{}
	lines := strings.Split(input, "\n")
	maxX := len(lines[0])
	maxY := len(lines) - 1
	startId := 0

	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		for x, char := range line {
			id := getId(maxX, x, y)
			neighbors := make([]int, 0)
			if char != 'S' {
				for _, pos := range TileMapping[char] {
					nx, ny := x+pos.x, y+pos.y
					if clamp(nx, 0, maxX) != nx || clamp(ny, 0, maxY) != ny {
						continue
					}
					neighbors = append(neighbors, getId(maxX, nx, ny))
				}
			}
			pipes[id] = pipe{
				id:        id,
				tile:      char,
				neighbors: neighbors,
			}
			if char == 'S' {
				startId = id
			}
		}
	}

	startNeighbors := make([]int, 0)
	for _, delta := range []position{
		{x: 0, y: -1},
		{x: 0, y: 1},
		{x: 1, y: 0},
		{x: -1, y: 0},
	} {
		idx := startId + delta.x + maxX*delta.y

		if slices.Contains(pipes[idx].neighbors, startId) {
			startNeighbors = append(startNeighbors, idx)
		}
	}
	pipes[startId] = pipe{
		id:        startId,
		tile:      'S',
		neighbors: startNeighbors,
	}

	return pipeField{
		start: startId,
		x:     maxX,
		y:     maxY,
		pipes: pipes,
	}
}

func day10part1(field pipeField) int {
	toCheck := []int{field.start}
	traversed := map[int]int{
		field.start: 0,
	}

	maximum := 0
	for len(toCheck) > 0 {
		idx := toCheck[0]
		pipe := field.pipes[idx]
		cost := traversed[idx]
		if maximum < cost {
			maximum = cost
		}
		for _, neighbor := range pipe.neighbors {
			if traversed[neighbor] == 0 || traversed[neighbor] > cost+1 {
				traversed[neighbor] = cost + 1
				toCheck = append(toCheck, neighbor)
			}
		}
		toCheck = toCheck[1:]
	}

	return maximum
}

func day10part2(field pipeField) int {
	// I realize now there's an infinitely better way to do this,
	// I'll get around to implementing it with my rust code

	traversed := map[int]bool{field.start: true}
	toCheck := []int{field.start}
	for len(toCheck) > 0 {
		idx := toCheck[0]
		pipe := field.pipes[idx]
		for _, neighbor := range pipe.neighbors {
			if !traversed[neighbor] {
				traversed[neighbor] = true
				toCheck = append(toCheck, neighbor)
			}
		}
		toCheck = toCheck[1:]
	}

	expanded := map[int]bool{}
	maxX := field.x * 3
	maxY := field.y * 3
	for id := range traversed {
		tile := field.pipes[id].tile
		x, y := id%field.x, id/field.x
		mx, my := 3*x+1, 3*y+1
		for _, p := range append([]position{{x: 0, y: 0}}, TileMapping[tile]...) {
			nx, ny := mx+p.x, my+p.y
			expanded[getId(maxX, nx, ny)] = true
		}
	}
	flooded := map[int]bool{0: true}
	toCheck = append(toCheck, 0)
	for len(toCheck) > 0 {
		id := toCheck[0]
		x, y := id%maxX, id/maxX
		for _, d := range []position{
			{x: -1, y: 0},
			{x: 0, y: -1},
			{x: 0, y: 1},
			{x: 1, y: 0},
		} {
			nx, ny := x+d.x, y+d.y
			if clamp(nx, 0, maxX) == nx && clamp(ny, 0, maxY) == ny {
				nId := getId(maxX, nx, ny)
				if !expanded[nId] && !flooded[nId] {
					flooded[nId] = true
					toCheck = append(toCheck, nId)
				}
			}
		}
		toCheck = toCheck[1:]
	}

	area := 0
	for y := 0; y < field.y; y++ {
		for x := 0; x < field.x; x++ {
			mx, my := 3*x+1, 3*y+1
			id := getId(maxX, mx, my)
			if !expanded[id] && !flooded[id] {
				area += 1
			}
		}
	}

	return area
}

func day10(part int, testing bool) {
	input := DayTenTestInput
	if !testing {
		input = readInput(10, false)
	} else if part == 2 {
		input = DayTenPartTwoTestInput
	}

	field := createPipeField(input)
	solver := day10part1
	s := "distance"
	if part == 2 {
		solver = day10part2
		s = "area"
	}

	fmt.Printf("The %s is %d\n", s, solver(field))
}
