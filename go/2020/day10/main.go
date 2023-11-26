package main

import (
	_ "embed"
	"fmt"
	"hora.dev/aoc/collection"
	"slices"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func main() {
	part1(input)
	//part2(example)
}

func part1(input string) {
	joltages := parse(input)
	slices.Sort(joltages)

	diffs := map[int]int{}
	current := 0

	collection.Each(joltages, func(joltage int) {
		diff := joltage - current
		if 1 <= diff && diff <= 3 {
			diffs[diff] += 1
			current = joltage
		}
	})

	diffs[3] += 1

	fmt.Printf("part 1: %d\n", diffs[1]*diffs[3])
}

func part2(input string) {
	joltages := parse(input)
	slices.Sort(joltages)

	current := 0
	possiblePaths := 0
	for i := 0; i < len(joltages); i++ {
		joltage := joltages[i]

		if findPaths(joltages[i:], current) {
			possiblePaths += 1
		}
		current = joltage
	}
	fmt.Printf("part 2: %d\n", possiblePaths)
}

func findPaths(adapters []int, current int) bool {
	for i := 0; i < len(adapters); i++ {
		joltage := adapters[i]
		currentDiff := joltage - current
		if !(currentDiff >= 1 && currentDiff <= 3) {
			return false // path not possible
		}
		current = joltage
	}
	return true
}

func parse(input string) []int {
	lines := strings.Split(input, "\n")
	nums, err := collection.MapErr(lines, strconv.Atoi)
	if err != nil {
		panic(err.Error())
	}
	return nums
}
