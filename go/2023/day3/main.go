package main

import (
	_ "embed"
	"fmt"
	"golang.org/x/exp/maps"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	schematics := strings.Split(input, "\n")
	width := len(schematics[0])
	result := 0

	for row, line := range schematics {
		acc := ""
		for col, char := range line {
			if unicode.IsDigit(char) {
				acc += string(char)
			} else if acc != "" {
				if isPartNumber(schematics, row, col-len(acc), len(acc)) {
					n, _ := strconv.Atoi(acc)
					result += n
				}
				acc = ""
			}
		}

		// anything left in the end of the row?
		if acc != "" && isPartNumber(schematics, row, width-len(acc), len(acc)) {
			n, _ := strconv.Atoi(acc)
			result += n
		}
	}

	fmt.Printf("part 1: %d\n", result)
}

func part2(input string) {
	schematics := strings.Split(input, "\n")
	result := 0

	for r, line := range schematics {
		for c, char := range line {
			if char == '*' {
				adj := findAdjacentNumbers(schematics, r, c)
				if len(adj) == 2 {
					ratio := 1
					for _, num := range adj {
						ratio *= num
					}
					result += ratio
				}
			}
		}
	}

	fmt.Printf("part 2: %d\n", result)
}

func findAdjacentNumbers(schematics []string, row, col int) []int {
	width := len(schematics[0])
	height := len(schematics)

	minRow := max(0, row-1)
	maxRow := min(height-1, row+1)
	minCol := max(0, col-1)
	maxCol := min(width-1, col+1)

	adjacents := map[int]struct{}{}

	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			char := rune(schematics[r][c])
			if unicode.IsDigit(char) {
				n := findNumberAt(schematics, r, c)
				adjacents[n] = struct{}{}
			}
		}
	}

	return maps.Keys(adjacents)
}

func findNumberAt(schematics []string, row, col int) int {
	thisRow := schematics[row]
	// find start
	start := 0
	for c := col; c >= 0; c-- {
		if !unicode.IsDigit(rune(thisRow[c])) {
			start = c + 1
			break
		}
	}
	// find end
	end := len(thisRow)
	for c := col; c < len(thisRow); c++ {
		if !unicode.IsDigit(rune(thisRow[c])) {
			end = c
			break
		}
	}
	n, _ := strconv.Atoi(thisRow[start:end])
	return n
}

func isPartNumber(schematics []string, row, col, size int) bool {
	width := len(schematics[0])
	height := len(schematics)

	minRow := max(0, row-1)
	maxRow := min(height-1, row+1)
	minCol := max(0, col-1)
	maxCol := min(width-1, col+size)

	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			char := rune(schematics[r][c])
			if isSymbol(char) {
				return true
			}
		}
	}

	return false
}

func isSymbol(char rune) bool {
	return !(unicode.IsDigit(char) || char == '.')
}
