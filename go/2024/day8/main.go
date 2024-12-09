package main

import (
	_ "embed"
	"fmt"
	"strings"

	"hora.dev/aoc/utils"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	rows := len(lines)
	cols := len(lines[0])

	antennas := parse(input)
	projections := map[[2]int]struct{}{}
	for _, pos := range antennas {
		castProjectionsPart1(pos, projections, rows, cols)
	}

	fmt.Printf("part 1: %v\n", len(projections))
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	rows := len(lines)
	cols := len(lines[0])

	antennas := parse(input)

	projections := map[[2]int]struct{}{}
	for _, pos := range antennas {
		castProjectionsPart2(pos, projections, rows, cols)
	}

	fmt.Printf("part 2: %v\n", len(projections))
}

func castProjectionsPart1(antennas [][2]int, dst map[[2]int]struct{}, maxRow, maxCol int) {
	for i, a1 := range antennas {
		for _, a2 := range antennas[i+1:] {
			p1, p2 := projectTwo(a1, a2, 1)
			p1Inside := p1[0] >= 0 && p1[0] < maxRow && p1[1] >= 0 && p1[1] < maxCol
			p2Inside := p2[0] >= 0 && p2[0] < maxRow && p2[1] >= 0 && p2[1] < maxCol
			if p1Inside {
				dst[p1] = struct{}{}
			}
			if p2Inside {
				dst[p2] = struct{}{}
			}
		}
	}
}

func castProjectionsPart2(antennas [][2]int, dst map[[2]int]struct{}, maxRow, maxCol int) {
	for i, a1 := range antennas {
		for _, a2 := range antennas[i+1:] {
			for _, p := range projectAll(a1, a2, maxRow, maxCol) {
				dst[p] = struct{}{}
			}
		}
	}
}

func projectAll(a1, a2 [2]int, rows, cols int) [][2]int {
	projections := [][2]int{a1, a2}
	multiplier := 0

	for {
		multiplier += 1
		p1, p2 := projectTwo(a1, a2, multiplier)

		p1Outisde := p1[0] < 0 || p1[0] >= rows || p1[1] < 0 || p1[1] >= cols
		p2Outside := p2[0] < 0 || p2[0] >= rows || p2[1] < 0 || p2[1] >= cols

		if !p1Outisde {
			projections = append(projections, p1)
		}
		if !p2Outside {
			projections = append(projections, p2)
		}
		if p1Outisde && p2Outside {
			break
		}
	}

	return projections
}

func projectTwo(p1, p2 [2]int, multiplier int) ([2]int, [2]int) {
	row1, col1 := p1[0], p1[1]
	row2, col2 := p2[0], p2[1]

	rowDiff := utils.Abs(p1[0], p2[0]) * multiplier
	colDiff := utils.Abs(p1[1], p2[1]) * multiplier

	var proj1, proj2 [2]int
	if row1 < row2 {
		// p1 is on top of p2
		if col1 < col2 {
			// p1 is on the left of p2
			proj1 = [2]int{row1 - rowDiff, col1 - colDiff}
			proj2 = [2]int{row2 + rowDiff, col2 + colDiff}
		} else {
			// p1 is on the right
			proj1 = [2]int{row1 - rowDiff, col1 + colDiff}
			proj2 = [2]int{row2 + rowDiff, col2 - colDiff}
		}
	} else {
		// p1 is bellow
		if col1 < col2 {
			// p1 is on the left
			proj1 = [2]int{row1 + rowDiff, col1 - colDiff}
			proj2 = [2]int{row2 - rowDiff, col2 + colDiff}
		} else {
			// p1 is on the right
			proj1 = [2]int{row1 + rowDiff, col1 + colDiff}
			proj2 = [2]int{row2 - rowDiff, col2 - colDiff}
		}
	}
	return proj1, proj2
}

func parse(input string) map[string][][2]int {
	res := map[string][][2]int{}
	for row, line := range strings.Split(input, "\n") {
		for col, char := range strings.Split(line, "") {
			if char == "." {
				continue
			}
			if pos, ok := res[char]; ok {
				res[char] = append(pos, [2]int{row, col})
			} else {
				res[char] = [][2]int{{row, col}}
			}
		}
	}
	return res
}
