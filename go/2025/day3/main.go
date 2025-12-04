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
	var res int64
	for l := range strings.Lines(input) {
		res += findMax(l, 2)
	}
	fmt.Printf("Part 1: %v\n", res)
}

func part2(input string) {
	var res int64
	for l := range strings.Lines(input) {
		res += findMax(l, 12)
	}
	fmt.Printf("Part 2: %v\n", res)
}

func findMax(l string, size int) int64 {
	ints := utils.IntList(l, "")
	pos := 0
	found := 0
	acc := int64(0)
	for {
		m := int64(0)
		mp := pos
		for i := pos; i <= len(ints)-size+found; i++ {
			if ints[i] > m {
				m = ints[i]
				mp = i
			}
		}
		acc = acc*10 + m
		pos = mp + 1
		found++
		if found == size {
			return acc
		}
	}
}
