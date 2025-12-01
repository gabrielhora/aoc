package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
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
	const size = 100
	pos := 50
	res := 0
	for l := range strings.Lines(input) {
		l = strings.TrimSpace(l)
		c, _ := strconv.Atoi(l[1:])
		if strings.HasPrefix(l, "L") {
			pos = pos - (c % size)
		} else {
			pos = pos + (c % size)
		}
		if pos < 0 {
			pos = size - (pos * -1)
		} else if pos > 99 {
			pos = pos - size
		}
		if pos == 0 {
			res++
		}
	}
	fmt.Printf("part 1: %v\n", res)
}

func part2(input string) {
	const size = 100
	pos := 50
	res := 0
	for l := range strings.Lines(input) {
		l = strings.TrimSpace(l)
		c, _ := strconv.Atoi(l[1:])
		dir := 1
		if strings.HasPrefix(l, "L") {
			dir = -1
		}
		for range c {
			pos = (pos + dir) % size
			if pos < 0 {
				pos = size - (pos * -1)
			} else if pos >= size {
				pos = pos - size
			}
			if pos == 0 {
				res++
			}
		}
	}
	fmt.Printf("part 2: %v\n", res)
}
