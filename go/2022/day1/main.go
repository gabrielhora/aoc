package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	part1(lines)
	part2(lines)
}

func part2(lines []string) {
	var counts []int
	sum := 0
	for _, line := range lines {
		if line == "" {
			counts = append(counts, sum)
			sum = 0
		}
		c, _ := strconv.Atoi(line)
		sum += c
	}
	counts = append(counts, sum)
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	result := 0
	for _, top := range counts[:3] {
		result += top
	}
	fmt.Printf("Part 2: %d\n", result)
}

func part1(lines []string) {
	biggest := 0
	sum := 0
	for _, line := range lines {
		if line == "" {
			if sum > biggest {
				biggest = sum
			}
			sum = 0
		}
		c, _ := strconv.Atoi(line)
		sum += c
	}
	fmt.Printf("Part 1: %d\n", biggest)
}
