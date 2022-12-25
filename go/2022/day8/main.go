package main

import (
	_ "embed"
	"fmt"
	"hora.dev/aoc/2022/utils/slice"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	grid := parse(input)
	size := len(grid[0])

	ans := 0
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if isVisible(grid, x, y) {
				ans += 1
			}
		}
	}

	fmt.Printf("part 1: %d\n", ans)
}

func part2(input string) {
	grid := parse(input)
	size := len(grid[0])

	best := float64(0)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			s := computeScoreAt(grid, x, y)
			best = math.Max(best, float64(s))
		}
	}

	fmt.Printf("part 2: %d\n", int(best))
}

func isVisible(grid [][]int, x int, y int) bool {
	size := len(grid[0])

	// edges are always visible
	if x == size-1 || x == 0 || y == size-1 || y == 0 {
		return true
	}

	// row and column for the x, y position
	height := grid[x][y]
	row := grid[x]
	var col []int
	for r := 0; r < size; r++ {
		col = append(col, grid[r][y])
	}

	left, right := row[:y], row[y+1:]
	up, down := col[:x], col[x+1:]

	return isVisibleFrom(left, height) ||
		isVisibleFrom(right, height) ||
		isVisibleFrom(up, height) ||
		isVisibleFrom(down, height)
}

func score(trees []int, height int) int {
	ans := 0
	for _, t := range trees {
		if t >= height {
			ans += 1
			break
		}
		ans += 1
	}
	return ans
}

func computeScoreAt(grid [][]int, x int, y int) int {
	size := len(grid[0])
	if x == size-1 || x == 0 || y == size-1 || y == 0 {
		return 0
	}

	height := grid[x][y]
	row := grid[x]
	var col []int
	for r := 0; r < size; r++ {
		col = append(col, grid[r][y])
	}
	left, right := row[:y], row[y+1:]
	up, down := col[:x], col[x+1:]

	// reverse left and up
	left = slice.Reverse(left)
	up = slice.Reverse(up)

	return score(left, height) *
		score(right, height) *
		score(up, height) *
		score(down, height)
}

func isVisibleFrom(trees []int, height int) bool {
	for _, t := range trees {
		if t >= height {
			return false
		}
	}
	return true
}

func parse(input string) [][]int {
	var lines [][]int
	for _, l := range strings.Split(input, "\n") {
		var chars []int
		for _, cs := range strings.Split(l, "") {
			c, _ := strconv.Atoi(cs)
			chars = append(chars, c)
		}
		lines = append(lines, chars)
	}
	return lines
}
