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
}

func part1(input string) {
	gardens := parse(input)
	rows := len(gardens)
	cols := len(gardens[0])

	// positions that were already fenced
	fencedPositions := map[[2]int]bool{}
	totalFences := 0
	totalArea := 0
	totalCost := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if _, ok := fencedPositions[[2]int{row, col}]; ok {
				continue
			}
			f, a := findArea(gardens, row, col, &fencedPositions)
			fmt.Printf("%s  %v * %v = %v\n", string(gardens[row][col]), a, f, f*a)
			totalFences += f
			totalArea += a
			totalCost += f * a
		}
	}

	fmt.Printf("part 1: %v\n", totalCost)
}

func findArea(gardens [][]rune, row, col int, fenced *map[[2]int]bool) (fences int, area int) {
	if _, ok := (*fenced)[[2]int{row, col}]; ok {
		return 0, 0
	}
	(*fenced)[[2]int{row, col}] = true

	garden := gardens[row][col]
	rows := len(gardens)
	cols := len(gardens[0])

	area = 1

	// edges of the map
	if col == 0 || col == cols-1 {
		fences += 1
	}
	if row == 0 || row == rows-1 {
		fences += 1
	}

	// up
	if row > 0 {
		if gardens[row-1][col] == garden {
			ff, aa := findArea(gardens, row-1, col, fenced)
			fences += ff
			area += aa
		} else {
			fences += 1
		}
	}

	// left
	if col > 0 {
		if gardens[row][col-1] == garden {
			ff, aa := findArea(gardens, row, col-1, fenced)
			fences += ff
			area += aa
		} else {
			fences += 1
		}
	}

	// right
	if col < cols-1 {
		if gardens[row][col+1] == garden {
			ff, aa := findArea(gardens, row, col+1, fenced)
			fences += ff
			area += aa
		} else {
			fences += 1
		}
	}

	// down
	if row < rows-1 {
		if gardens[row+1][col] == garden {
			ff, aa := findArea(gardens, row+1, col, fenced)
			fences += ff
			area += aa
		} else {
			fences += 1
		}
	}

	return
}

func parse(input string) [][]rune {
	var res [][]rune
	for _, row := range strings.Split(input, "\n") {
		res = append(res, []rune(row))
	}
	return res
}
