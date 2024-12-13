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
	tmap := parse(input)

	ans := 0
	for _, start := range findTrailheads(tmap) {
		var paths [][][2]int
		walkPath(tmap, start, [][2]int{start}, &paths)

		uniqueDestinations := map[[2]int]bool{}
		for _, p := range paths {
			uniqueDestinations[p[len(p)-1]] = true
		}
		ans += len(uniqueDestinations)
	}

	fmt.Printf("part 1: %v\n", ans)
}

func part2(input string) {
	tmap := parse(input)

	ans := 0
	for _, start := range findTrailheads(tmap) {
		var paths [][][2]int
		walkPath(tmap, start, [][2]int{start}, &paths)
		ans += len(paths)
	}

	fmt.Printf("part 2: %v\n", ans)
}

var dirs = [][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func walkPath(tmap [][]int64, pos [2]int, path [][2]int, acc *[][][2]int) {
	curValue := tmap[pos[0]][pos[1]]
	if curValue == 9 { // base case
		copyOfPath := append([][2]int{}, path...)
		*acc = append(*acc, copyOfPath)
		return
	}

	// compute possible moves
	var moves [][2]int
	for _, dir := range dirs {
		nextPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		if canMove(tmap, nextPos, curValue) {
			moves = append(moves, nextPos)
		}
	}

	// check possible paths
	for _, move := range moves {
		path = append(path, move)
		walkPath(tmap, move, path, acc)
		path = path[:len(path)-1] // backtrack
	}
}

func checkBounds(tmap [][]int64, pos [2]int) bool {
	return pos[0] >= 0 && pos[0] < len(tmap) && pos[1] >= 0 && pos[1] < len(tmap[1])
}

func canMove(tmap [][]int64, nextPos [2]int, curValue int64) bool {
	if !checkBounds(tmap, nextPos) {
		return false
	}
	nextVal := tmap[nextPos[0]][nextPos[1]]
	return nextVal-curValue == 1 // can move there if change by +1
}

func findTrailheads(tmap [][]int64) [][2]int {
	var starts [][2]int
	for r, row := range tmap {
		for c, col := range row {
			if col == 0 {
				starts = append(starts, [2]int{r, c})
			}
		}
	}
	return starts
}

func parse(input string) [][]int64 {
	var result [][]int64
	for _, line := range strings.Split(input, "\n") {
		result = append(result, utils.IntList(line, ""))
	}
	return result
}
