package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
}

func part1(input string) {
	parts := strings.Split(input, "\n\n")
	grid := parseGrid(parts[0])
	moves := parseMoves(parts[1])

	startCol := findFirstNonVoid(grid[1])
	row, col, dir := move(grid, moves, [2]int{1, startCol}, right)

	ans := 1000*row + 4*col + dir.value()
	fmt.Printf("part 1: %d\n", ans)
}

func printGrid(grid [][]block, pos [2]int, dir direction) {
	fmt.Print("\033[H\033[2J") // clear the window

	for r, row := range grid {
		for c, col := range row {
			if r == pos[0] && c == pos[1] {
				fmt.Print(dir)
			} else if col == open {
				fmt.Print(".")
			} else if col == wall {
				fmt.Print("#")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
	//time.Sleep(1 * time.Second)
}

func move(grid [][]block, moves [][]direction, start [2]int, startDir direction) (row, col int, facing direction) {
	pos := start
	curDir := startDir

	for _, moveSet := range moves {
		for _, move := range moveSet {
			curDir = move
			nextPos := [2]int{pos[0] + move[0], pos[1] + move[1]}
			nextBlock := grid[nextPos[0]][nextPos[1]]
			if nextBlock == wall {
				// if hit a wall go to next moveset
				break
			} else if nextBlock == void {
				// circle back to the begining (depends on curDir)
				// begining can also be a wall, in which case move to next moveset
				newPos, isWall := circleBack(grid, pos, curDir)
				if isWall {
					break
				}
				pos = newPos
			} else {
				// otherwise move to the new position
				pos = nextPos
			}

			//printGrid(grid, pos, curDir)
		}
	}

	return pos[0], pos[1], curDir
}

// circleBack returns the new position at the begining of the row/col, if that position
// is a wall it returns false
func circleBack(grid [][]block, pos [2]int, dir direction) ([2]int, bool) {
	if dir == right {
		col := findFirstNonVoid(grid[pos[0]])
		if grid[pos[0]][col] == wall {
			return [2]int{}, true
		}
		return [2]int{pos[0], col}, false
	}

	if dir == left {
		col := findLastNonVoid(grid[pos[0]])
		if grid[pos[0]][col] == wall {
			return [2]int{}, true
		}
		return [2]int{pos[0], col}, false
	}

	if dir == up || dir == down {
		var col []block
		for row := 0; row < len(grid); row++ {
			col = append(col, grid[row][pos[1]])
		}

		if dir == up {
			row := findLastNonVoid(col)
			if grid[row][pos[1]] == wall {
				return [2]int{}, true
			}
			return [2]int{row, pos[1]}, false
		}
		if dir == down {
			row := findFirstNonVoid(col)
			if grid[row][pos[1]] == wall {
				return [2]int{}, true
			}
			return [2]int{row, pos[1]}, false
		}
	}

	panic("invalid direction")
}

func findFirstNonVoid(list []block) int {
	for i := 0; i < len(list); i++ {
		if list[i] == wall || list[i] == open {
			return i
		}
	}
	return -1
}

func findLastNonVoid(list []block) int {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == wall || list[i] == open {
			return i
		}
	}
	return -1
}

func parseGrid(grid string) [][]block {
	// grid == [row][col]block

	// find biggest row
	biggestRow := 0
	for _, row := range strings.Split(grid, "\n") {
		if len(row) > biggestRow {
			biggestRow = len(row)
		}
	}

	// the first and last rows are all void
	var allVoidRow []block
	for i := 0; i < biggestRow+2; i++ {
		allVoidRow = append(allVoidRow, void)
	}

	blocks := [][]block{allVoidRow}

	for _, row := range strings.Split(grid, "\n") {
		// start with an initial void column
		rowBlocks := []block{void}
		for col := 0; col < biggestRow; col++ {
			if col >= len(row) {
				rowBlocks = append(rowBlocks, void)
			} else if row[col] == '#' {
				rowBlocks = append(rowBlocks, wall)
			} else if row[col] == '.' {
				rowBlocks = append(rowBlocks, open)
			} else {
				rowBlocks = append(rowBlocks, void)
			}
		}
		// add the last column as void
		rowBlocks = append(rowBlocks, void)
		blocks = append(blocks, rowBlocks)
	}

	// add the last row as all void
	blocks = append(blocks, allVoidRow)
	return blocks
}

// parseMoves returns a slice of slices of direction deltas, so a string "2R3" will return:
//
//	[
//	  [[1 0] [1 0]], // two right movements then turn clockwise
//	  [[0 1] [0 1] [0 1]], // three down movements
//	]
func parseMoves(input string) [][]direction {
	var parts []string
	acc := ""
	for _, c := range input {
		if c == 'R' {
			parts = append(parts, acc)
			parts = append(parts, "R")
			acc = ""
		} else if c == 'L' {
			parts = append(parts, acc)
			parts = append(parts, "L")
			acc = ""
		} else {
			acc += string(c)
		}
	}
	if acc != "" {
		parts = append(parts, acc)
	}

	var moves [][]direction
	cur := right
	for _, m := range parts {
		if m == "R" {
			cur = cur.turnClockwise()
		} else if m == "L" {
			cur = cur.turnCounterclockwise()
		} else {
			n, _ := strconv.Atoi(m)
			var dirMoves []direction
			for i := 0; i < n; i++ {
				dirMoves = append(dirMoves, cur)
			}
			moves = append(moves, dirMoves)
		}
	}

	return moves
}

type block int

const (
	void block = iota
	open
	wall
)

type direction [2]int

var (
	// row, col
	right = direction{0, 1}
	down  = direction{1, 0}
	left  = direction{0, -1}
	up    = direction{-1, 0}
)

func (d direction) String() string {
	switch d {
	case right:
		return ">"
	case down:
		return "v"
	case left:
		return "<"
	case up:
		return "^"
	default:
		panic("invalid direction")
	}
}

func (d direction) turnClockwise() direction {
	switch d {
	case right:
		return down
	case down:
		return left
	case left:
		return up
	case up:
		return right
	default:
		panic("invalid direction")
	}
}

func (d direction) turnCounterclockwise() direction {
	switch d {
	case right:
		return up
	case down:
		return right
	case left:
		return down
	case up:
		return left
	default:
		panic("invalid direction")
	}
}

func (d direction) value() int {
	switch d {
	case right:
		return 0
	case down:
		return 1
	case left:
		return 2
	case up:
		return 3
	default:
		panic("invalid direction")
	}
}
