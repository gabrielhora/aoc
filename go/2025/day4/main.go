package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
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
	grid := parse(input)
	res := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if count(grid, r, c) < 4 {
				res++
			}
		}
	}
	fmt.Printf("Part 1: %v\n", res)
}

func part2(input string) {
	grid := parse(input)
	res := 0
	for {
		var changes [][2]int
		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[r]); c++ {
				if count(grid, r, c) < 4 {
					changes = append(changes, [2]int{r, c})
					res++
				}
			}
		}
		if len(changes) == 0 {
			break
		}
		for _, cs := range changes {
			grid[cs[0]][cs[1]] = "."
		}
		changes = nil
	}
	fmt.Printf("Part 2: %v\n", res)
}

func count(grid [][]string, row, col int) int {
	if grid[row][col] != "@" {
		return math.MaxInt
	}
	cnt := 0
	for r := row - 1; r <= row+1; r++ {
		if r < 0 || r >= len(grid) {
			continue
		}
		for c := col - 1; c <= col+1; c++ {
			if (r == row && c == col) || c < 0 || c >= len(grid[r]) {
				continue
			}
			if grid[r][c] == "@" {
				cnt++
			}
		}
	}
	return cnt
}

func parse(input string) [][]string {
	var res [][]string
	for l := range strings.Lines(input) {
		r := strings.Split(strings.TrimSpace(l), "")
		res = append(res, r)
	}
	return res
}
