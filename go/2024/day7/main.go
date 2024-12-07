package main

import (
	_ "embed"
	"fmt"
	"hora.dev/aoc/utils"
	"slices"
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
	ans := int64(0)
	for _, eq := range parseEquations(input) {
		var results []int64
		resolveAll(eq.values, []string{"+", "*"}, eq.values[0], 0, &results)
		if slices.Contains(results, eq.target) {
			ans += eq.target
		}
	}
	fmt.Printf("part 1: %v\n", ans)
}

func part2(input string) {
	ans := int64(0)
	for _, eq := range parseEquations(input) {
		var results []int64
		resolveAll(eq.values, []string{"+", "*", "||"}, eq.values[0], 0, &results)
		if slices.Contains(results, eq.target) {
			ans += eq.target
		}
	}
	fmt.Printf("part 1: %v\n", ans)
}

func resolveAll(vals []int64, ops []string, cur int64, idx int, results *[]int64) {
	if idx == len(vals)-1 {
		// all consts were consumed, add result of equation and return
		*results = append(*results, cur)
		return
	}

	// calculate the rest of the equation by applying each possible operator
	for _, op := range ops {
		nextConst := vals[idx+1]
		var newCur int64
		switch op {
		case "+":
			newCur = cur + nextConst
		case "*":
			newCur = cur * nextConst
		case "||":
			newCur = utils.ToInt64(fmt.Sprintf("%d%d", cur, nextConst))
		default:
			panic("invalid operator")
		}

		resolveAll(vals, ops, newCur, idx+1, results)
	}
}

type equation struct {
	target int64
	values []int64
}

func parseEquations(input string) []equation {
	var equations []equation
	for _, line := range strings.Split(input, "\n") {
		ans := utils.ToInt64(line[:strings.Index(line, ":")])
		consts := utils.IntList(line[strings.Index(line, ":")+1:], " ")
		equations = append(equations, equation{ans, consts})
	}
	return equations
}
