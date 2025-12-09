package main

import (
	_ "embed"
	"fmt"
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
	grid, row, _ := parse(input)
	splits := 0
	for i := row + 1; i < len(grid); i++ {
		splits += countSplits(i, grid)
	}
	fmt.Printf("Part 1: %v\n", splits)
}

func part2(input string) {
	grid, row, col := parse(input)
	paths := countPaths(grid, row, col)
	fmt.Printf("Part 2: %v\n", paths)
}

func parse(input string) ([][]string, int, int) {
	var grid [][]string
	var startRow, startCol int
	for r, l := range strings.Split(input, "\n") {
		if strings.Contains(l, "S") {
			startRow = r
			startCol = strings.Index(l, "S")
		}
		grid = append(grid, strings.Split(strings.TrimSpace(l), ""))
	}
	return grid, startRow, startCol
}

func countSplits(row int, grid [][]string) int {
	splits := 0
	for col := 0; col < len(grid[row]); col++ {
		top := grid[row-1][col]
		cur := grid[row][col]
		prev := ""
		if col > 0 {
			prev = grid[row][col-1]
		}
		next := ""
		if col < len(grid[row])-1 {
			next = grid[row][col+1]
		}
		switch {
		case cur == "^" && top == "|":
			cur = "^"
			splits++
		case prev == "^" || next == "^" || top == "|" || top == "S":
			cur = "|"
		}
		grid[row][col] = cur
	}
	return splits
}

var seen = map[[2]int]int64{}

func visit(grid [][]string, row, col int) int64 {
	res := countPaths(grid, row, col)
	seen[[2]int{row, col}] = res
	return res
}

func countPaths(grid [][]string, row, col int) int64 {
	if v, ok := seen[[2]int{row, col}]; ok {
		return v
	}
	if row >= len(grid) {
		seen[[2]int{row, col}] = 1
		return 1
	}
	if grid[row][col] == "." || grid[row][col] == "S" {
		return visit(grid, row+1, col)
	} else if grid[row][col] == "^" {
		return visit(grid, row, col-1) + visit(grid, row, col+1)
	}
	panic("not possible")
}
