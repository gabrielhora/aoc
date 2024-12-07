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
	//part1(input)
	part2(input) // 1455 too low
}

func part1(input string) {
	grid := utils.GridStr(input, "\n", "")
	guard, dir := findGuard(grid)
	blocks := findBlocks(grid)
	visited, _ := moveGuard(guard, dir, grid, blocks, false)
	fmt.Printf("part 1: %d\n", len(visited))
}

func part2(input string) {
	grid := utils.GridStr(input, "\n", "")
	guard, dir := findGuard(grid)
	blocks := findBlocks(grid)

	possibleNewBlockPos, _ := moveGuard(guard, dir, grid, blocks, false)
	blocks = append(blocks, [2]int{0, 0})
	ans := 0
	// try to put new blocks (index 0 is the guard position)
	for i, newBlockPos := range possibleNewBlockPos[1:] {
		if i%100 == 0 {
			fmt.Printf("checking %d of %d\n", i, len(possibleNewBlockPos)-1)
		}
		blocks[len(blocks)-1] = newBlockPos
		if _, loops := moveGuard(guard, dir, grid, blocks, true); loops {
			ans += 1
		}
	}
	fmt.Printf("part 2: %d\n", ans)
}

func moveGuard(start [2]int, dir [2]int, grid [][]string, blocks [][2]int, detectLoops bool) ([][2]int, bool) {
	rows, cols := len(grid), len(grid[0])
	currentPos := start
	var visited [][2]int
	for {
		//printGrid(rows, cols, currentPos, dir, visited, blocks)

		if detectLoops || !slices.Contains(visited, currentPos) {
			visited = append(visited, currentPos)
		}
		if detectLoops && detectLoop(visited) {
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
			currentPos = [2]int{currentPos[0] + dir[0], currentPos[1] + dir[1]}
			continue
		}
		// otherwise move there
		currentPos = newPos
	}
	return visited, false
}

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

func detectLoop2(visited [][2]int) bool {
	if len(visited) < 4 {
		return false
	}
	for i := 2; i < len(visited)/2+1; i++ {
		s1 := visited[len(visited)-i:]
		s2 := visited[len(visited)-i*2 : len(visited)-i]
		if utils.SlicesEqual(s1, s2) {
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
