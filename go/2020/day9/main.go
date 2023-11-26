package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func main() {
	part1Result := part1(input)
	fmt.Printf("part 1: %d\n", part1Result)
	part2(input, part1Result)
}

func part1(input string) int {
	var result int
	nums := parse(input)
	preambleSize := 25
	for i := preambleSize; i < len(nums); i++ {
		n := nums[i]
		prev := nums[i-preambleSize : i]
		if !findInSums(prev, n) {
			result = n
			break
		}
	}
	return result
}

func findInSums(coll []int, lookup int) bool {
	for i := 0; i < len(coll); i++ {
		for j := 0; j < len(coll); j++ {
			sum := coll[i] + coll[j]
			if sum == lookup {
				return true
			}
		}
	}
	return false
}

func part2(input string, target int) {
	nums := parse(input)
	sumSet := findSumSet(nums, target)
	smallest, largest := slices.Min(sumSet), slices.Max(sumSet)
	fmt.Printf("part 2: %d\n", smallest+largest)
}

func findSumSet(coll []int, lookup int) []int {
	for i := 0; i < len(coll); i++ {
		sum := coll[i]
		set := []int{coll[i]}

		for j := i + 1; j < len(coll); j++ {
			sum += coll[j]
			set = append(set, coll[j])
			if sum == lookup {
				return set
			}
			if sum > lookup {
				break
			}
		}
	}
	return nil
}

func parse(input string) []int {
	var nums []int
	for _, line := range strings.Split(input, "\n") {
		n, _ := strconv.Atoi(line)
		nums = append(nums, n)
	}
	return nums
}
