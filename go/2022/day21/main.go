package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	monkeys := parse(input)
	fmt.Printf("part 1: %d\n", solve(monkeys, monkeys["root"]))
}

func part2(input string) {
	monkeys := parse(input)

	// find the side of the equation that is not affected by "humn" value
	human, noHuman := monkeys["root"].eq.m1, monkeys["root"].eq.m2
	if findHuman(monkeys, monkeys[noHuman]) {
		noHuman, human = human, noHuman
	}

	target := solve(monkeys, monkeys[noHuman])
	ans := binarySearch(monkeys, monkeys[human], 0, 1000000000000000, target)
	fmt.Printf("part 2: %d\n", ans)
}

func binarySearch(monkeys map[string]*monkey, start *monkey, left, right, target int) int {
	if left > right {
		return -1 // not found
	}

	val := (left + right) / 2
	monkeys["humn"].value = &val
	res := solve(monkeys, start)

	if res < target {
		return binarySearch(monkeys, start, left, val+1, target)
	} else if res > target {
		return binarySearch(monkeys, start, val-1, right, target)
	} else {
		return val
	}
}

func solve(monkeys map[string]*monkey, start *monkey) int {
	if start.value != nil {
		return *start.value
	}
	val1 := solve(monkeys, monkeys[start.eq.m1])
	val2 := solve(monkeys, monkeys[start.eq.m2])
	switch start.eq.op {
	case "+":
		return val1 + val2
	case "-":
		return val1 - val2
	case "*":
		return val1 * val2
	case "/":
		return val1 / val2
	default:
		panic("invalid operation")
	}
}

func findHuman(monkeys map[string]*monkey, start *monkey) bool {
	if start.value != nil {
		return false
	}
	if start.eq.m1 == "humn" || start.eq.m2 == "humn" {
		return true
	}
	return findHuman(monkeys, monkeys[start.eq.m1]) ||
		findHuman(monkeys, monkeys[start.eq.m2])
}

var (
	equationRe = regexp.MustCompile(`(\w+): (\w+) (.) (\w+)`)
	valueRe    = regexp.MustCompile(`(\w+): (\d+)`)
)

type equation struct {
	m1, op, m2 string
}

type monkey struct {
	value *int
	eq    *equation
}

func parse(input string) map[string]*monkey {
	res := map[string]*monkey{}
	for _, l := range strings.Split(input, "\n") {
		m := &monkey{}
		if equationRe.MatchString(l) {
			caps := equationRe.FindStringSubmatch(l)
			m.eq = &equation{caps[2], caps[3], caps[4]}
			res[caps[1]] = m
		} else {
			caps := valueRe.FindStringSubmatch(l)
			val, _ := strconv.Atoi(caps[2])
			m.value = &val
			res[caps[1]] = m
		}
	}
	return res
}
