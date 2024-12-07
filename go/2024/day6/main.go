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
	part2(input)
}

func part1(input string) {
	grid := utils.GridStr(input, "\n", "")
	guard, dir := findGuard(grid)
	blocks := findBlocks(grid)
	visited, _ := moveGuard(guard, dir, grid, blocks, false)
	fmt.Printf("part 1: %d\n", len(visited))
}

func part2(input string) {
	fmt.Print("calculating part 2, takes a few seconds...\n")

	grid := utils.GridStr(input, "\n", "")
	guard, dir := findGuard(grid)
	blocks := findBlocks(grid)

	possibleNewBlockPos, _ := moveGuard(guard, dir, grid, blocks, false)
	possibleNewBlockPos = possibleNewBlockPos[1:] // index 0 is the guard pos
	blocks = append(blocks, [2]int{0, 0})

	ans := 0
	for i, newBlockPos := range possibleNewBlockPos {
		if i%1000 == 0 {
			fmt.Printf("checking %d of %d\n", i, len(possibleNewBlockPos)-1)
		}
		blocks[len(blocks)-1] = newBlockPos
		if _, isLooping := moveGuard(guard, dir, grid, blocks, true); isLooping {
			ans += 1
		}
	}
	fmt.Printf("part 2: %d\n", ans)
}

func moveGuard(start [2]int, dir [2]int, grid [][]string, blocks [][2]int, detectLoops bool) ([][2]int, bool) {
	rows, cols := len(grid), len(grid[0])
	currentPos := start
	var visited [][2]int
	moves := 0
	for {
		// print the grid with visited blocks for debugging
		//printGrid(rows, cols, currentPos, dir, visited, blocks)
		moves += 1

		if !detectLoops {
			if !slices.Contains(visited, currentPos) {
				visited = append(visited, currentPos)
			}
		} else if moves > 10000 {
			// this "stupid" loop detection is faster than actually looking for a loop
			// in the visited slice, so I'll leave it like this
			return visited, true
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
	return visited, false
}

// not used, just assuming a maxium number of movements is faster than this
// I'll leave it here for reference... maybe I'll figure out a better way
func detectLoop(visited [][2]int) bool {
	if len(visited) < 2 {
		return false
	}
	search := visited[len(visited)-2:]
	for i := len(visited) - 2; i >= 2; i-- {
		window := visited[i-2 : i]
		if utils.SlicesEqual(search, window) {
			return true
		}
	}
	return false
}

func findGuard(grid [][]string) ([2]int, [2]int) {
	for ri, r := range grid {
		for ci, c := range r {
			p := [2]int{ri, ci}
			if c == "^" {
				return p, [2]int{-1, 0}
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
