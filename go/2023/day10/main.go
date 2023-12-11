package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	m := parse(input)
	start := findStart(m)
	path, _ := findLoop(m, start)
	result := (len(path) + 1) / 2

	fmt.Printf("part 1: %d\n", result)
}

func part2(input string) {
	m := parse(input)
	start := findStart(m)
	path, vertices := findLoop(m, start)

	count := 0
	for y, row := range m {
		for x := range row {
			pt := point{x, y}
			if !slices.Contains(path, pt) && inside(pt, vertices) {
				count += 1
			}
		}
	}

	fmt.Printf("part 2: %v\n", count-1)
}

type point struct{ x, y int }

const (
	north = iota
	east
	south
	west
)

func walkMap(m [][]string, start point, dir int) ([]point, []point) {
	pos := start
	var path []point
	vertices := []point{start}

	for {
		//printMap(m, pos)

		switch dir {
		case north:
			pos.y -= 1
		case east:
			pos.x += 1
		case south:
			pos.y += 1
		case west:
			pos.x -= 1
		}

		if pos == start {
			return path, vertices // loop
		}
		if pos.y >= len(m) || pos.x >= len(m[0]) || pos.x < 0 || pos.y < 0 {
			return nil, nil // outside of map
		}

		switch m[int64(pos.y)][int64(pos.x)] {
		case "|":
			if !(dir == north || dir == south) {
				return nil, nil
			}
		case "-":
			if !(dir == east || dir == west) {
				return nil, nil
			}
		case "L":
			if dir == south {
				dir = east
				vertices = append(vertices, pos)
			} else if dir == west {
				dir = north
				vertices = append(vertices, pos)
			} else {
				return nil, nil
			}
		case "J":
			if dir == south {
				dir = west
				vertices = append(vertices, pos)
			} else if dir == east {
				dir = north
				vertices = append(vertices, pos)
			} else {
				return nil, nil
			}
		case "7":
			if dir == north {
				dir = west
				vertices = append(vertices, pos)
			} else if dir == east {
				dir = south
				vertices = append(vertices, pos)
			} else {
				return nil, nil
			}
		case "F":
			if dir == north {
				dir = east
				vertices = append(vertices, pos)
			} else if dir == west {
				dir = south
				vertices = append(vertices, pos)
			} else {
				return nil, nil
			}
		case ".":
			return nil, nil
		}

		path = append(path, pos)
	}
}

func findStart(m [][]string) point {
	for y, row := range m {
		for x, cell := range row {
			if cell == "S" {
				return point{x, y}
			}
		}
	}
	panic("should have a start")
}

func findLoop(m [][]string, start point) ([]point, []point) {
	for _, d := range []int{north, east, south, west} {
		if p, s := walkMap(m, start, d); p != nil {
			return p, s
		}
	}
	panic("should have at least one loop")
}

func parse(input string) [][]string {
	var result [][]string
	for _, l := range strings.Split(input, "\n") {
		result = append(result, strings.Split(l, ""))
	}
	return result
}

// https://rosettacode.org/wiki/Ray-casting_algorithm#Go
func inside(pt point, pg []point) bool {
	if len(pg) < 3 {
		return false
	}
	in := rayIntersectsSegment(pt, pg[len(pg)-1], pg[0])
	for i := 1; i < len(pg); i++ {
		if rayIntersectsSegment(pt, pg[i-1], pg[i]) {
			in = !in
		}
	}
	return in
}

func rayIntersectsSegment(p, a, b point) bool {
	return (a.y > p.y) != (b.y > p.y) &&
		p.x < (b.x-a.x)*(p.y-a.y)/(b.y-a.y)+a.x
}
