package main

import (
	_ "embed"
	"fmt"
	"hora.dev/aoc/2022/utils/set"
	"hora.dev/aoc/2022/utils/slice"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	chars := strings.Split(input, "")
	fmt.Printf("part 1: %d\n", firstUniqueSet(chars, 4))
}

func part2(input string) {
	chars := strings.Split(input, "")
	fmt.Printf("part 2: %d\n", firstUniqueSet(chars, 14))
}

func firstUniqueSet(chars []string, size int) int {
	for i, w := range slice.SlidingWindow(chars, size) {
		if len(set.FromValues(w...)) == size {
			return i + size
		}
	}
	return 0
}
