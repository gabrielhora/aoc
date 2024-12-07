package main

import (
	_ "embed"
	"fmt"
	"hora.dev/aoc/utils"
	"slices"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func main() {
	part1(input)
}

func part1(input string) {
	grid := utils.GridStr(input, "\n", "")
	guard, dir := findGuard(grid)
	blocks := findBlocks(grid)
	visited := moveGuard(guard, dir, grid, blocks)
	fmt.Printf("part 1: %d\n", len(visited))
}

func moveGuard(start [2]int, dir [2]int, grid [][]string, blocks [][2]int) [][2]int {
	rows, cols := len(grid), len(grid[0])
	currentPos := start
	var visited [][2]int
	for {
		//printGrid(rows, cols, currentPos, dir, visited, blocks)

		if !slices.Contains(visited, currentPos) {
			visited = append(visited, currentPos)
		}
		newPos := [2]int{currentPos[0] + dir[0], currentPos[1] + dir[1]}
		// if guard is outside, return
		if newPos[0] < 0 || newPos[0] >= rows || newPos[1] < 0 || newPos[1] >= cols {
			break
		}
		// if new pos is blocked
		if isBlocked(blocks, newPos) {
			dir = turn90Degrees(dir)
			continue
		}
		// otherwise move there
		currentPos = newPos
	}
	return visited
}

func findGuard(grid [][]string) ([2]int, [2]int) {
	for ri, r := range grid {
		for ci, c := range r {
			p := [2]int{ri, ci}
			if c == "^" {
				return p, [2]int{-1, 0}
			}
			if c == ">" {
				return p, [2]int{0, 1}
			}
			if c == "<" {
				return p, [2]int{0, -1}
			}
			if c == "v" {
				return p, [2]int{1, 0}
			}
		}
	}
	return [2]int{}, [2]int{}
}

func findBlocks(grid [][]string) [][2]int {
	var blocks [][2]int
	for ri, r := range grid {
		for ci, c := range r {
			if c == "#" {
				blocks = append(blocks, [2]int{ri, ci})
			}
		}
	}
	return blocks
}

func isBlocked(blocks [][2]int, pos [2]int) bool {
	for _, block := range blocks {
		if block == pos {
			return true
		}
	}
	return false
}

func turn90Degrees(dir [2]int) [2]int {
	// up -> right
	if dir[0] == -1 && dir[1] == 0 {
		return [2]int{0, 1}
	}
	// right -> down
	if dir[0] == 0 && dir[1] == 1 {
		return [2]int{1, 0}
	}
	// down -> left
	if dir[0] == 1 && dir[1] == 0 {
		return [2]int{0, -1}
	}
	// left -> up
	if dir[0] == 0 && dir[1] == -1 {
		return [2]int{-1, 0}
	}

	panic("not possible")
}

func printGrid(rows, cols int, guard [2]int, dir [2]int, visited [][2]int, blocks [][2]int) {
	for r := range rows {
		for c := range cols {
			p := [2]int{r, c}
			if p == guard {
				switch dir {
				case [2]int{-1, 0}:
					print("^")
				case [2]int{0, 1}:
					print(">")
				case [2]int{0, -1}:
					print("<")
				case [2]int{1, 0}:
					print("v")
				}
			} else if slices.Contains(blocks, p) {
				print("#")
			} else if slices.Contains(visited, p) {
				print("X")
			} else {
				print("_")
			}
		}
		print("\n")
	}
	print("\n")
}
