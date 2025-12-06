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
	res := solve(parsePart1(input))
	fmt.Printf("Part 1: %v\n", res)
}

func part2(input string) {
	res := solve(parsePart2(input))
	fmt.Printf("Part 2: %v\n", res)
}

func solve(eqs []equation) int64 {
	res := int64(0)
	for _, eq := range eqs {
		switch eq.op {
		case "+":
			for _, t := range eq.terms {
				res += t
			}
		case "*":
			er := int64(1)
			for _, t := range eq.terms {
				er *= t
			}
			res += er
		}
	}
	return res
}

type equation struct {
	op    string
	terms []int64
}

func parsePart1(input string) []equation {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	ops := utils.Filter(strings.Split(lines[len(lines)-1], ""), utils.IsBlank)
	lines = lines[:len(lines)-1]

	var nums [][]int64
	for _, l := range lines {
		cols := utils.IntListBySpaces[int64](l)
		nums = append(nums, cols)
	}

	var eqs []equation
	for col := 0; col < len(nums[0]); col++ {
		eq := equation{op: ops[0]}
		ops = ops[1:]
		for row := 0; row < len(nums); row++ {
			eq.terms = append(eq.terms, nums[row][col])
		}
		eqs = append(eqs, eq)
	}

	return eqs
}

func parsePart2(input string) []equation {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	ops := utils.Filter(strings.Split(lines[len(lines)-1], ""), utils.IsBlank)
	lines = lines[:len(lines)-1]

	var eqs []equation
	var terms []int64
	for c := 0; c < len(lines[0]); c++ {
		num := int64(0)
		for r := 0; r < len(lines); r++ {
			ch := lines[r][c]
			if ch == ' ' {
				continue
			}
			d := int64(ch - '0')
			num = num*10 + d
		}

		if num != 0 {
			terms = append(terms, num)
		} else {
			eqs = append(eqs, equation{
				op:    ops[0],
				terms: terms,
			})
			ops = ops[1:]
			terms = nil
		}
	}

	if len(terms) > 0 {
		eqs = append(eqs, equation{
			op:    ops[0],
			terms: terms,
		})
	}

	return eqs
}
