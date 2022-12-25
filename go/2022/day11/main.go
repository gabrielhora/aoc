package main

//10605
//2_713_310_158

import (
	_ "embed"
	"fmt"
	"math/big"
	"regexp"
	"sort"
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
	playRounds(monkeys, 20, func(i int) int { return i / 3 })

	var moves []int
	for _, m := range monkeys {
		moves = append(moves, m.moves)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(moves)))

	fmt.Printf("part 1: %d\n", moves[0]*moves[1])
}

func part2(input string) {
	monkeys := parse(input)

	div := monkeys[0].test
	for _, m := range monkeys {
		div = lcm(div, m.test)
	}

	playRounds(monkeys, 10000, func(i int) int { return i % div })
	var moves []int
	for _, m := range monkeys {
		moves = append(moves, m.moves)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(moves)))

	fmt.Printf("part 2: %d\n", moves[0]*moves[1])
}

type worryFn = func(int) int

func playRounds(monkeys []*monkey, rounds int, fn worryFn) {
	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			moveItems(monkeys, m, fn)
		}
	}
}

func moveItems(monkeys []*monkey, m *monkey, fn worryFn) {
	items := append([]int{}, m.items...)

	for _, item := range items {
		item = fn(runOp(m.op, item))

		target := monkeys[m.ifFalse]
		if item%m.test == 0 {
			target = monkeys[m.ifTrue]
		}

		// move the item to target
		target.items = append(target.items, item)
		m.items = m.items[1:]
		m.moves += 1
	}
}

func runOp(op string, item int) int {
	operation := op[4:5]
	op1 := item
	op2 := item
	if opRe.MatchString(op) {
		caps := opRe.FindStringSubmatch(op)
		op2, _ = strconv.Atoi(caps[1])
	}
	switch operation {
	case "+":
		return op1 + op2
	case "*":
		return op1 * op2
	default:
		panic("invalid operation")
	}
}

type monkey struct {
	items   []int
	op      string
	test    int
	ifTrue  int
	ifFalse int
	moves   int
}

var (
	monkeyRe = regexp.MustCompile(`Starting items: (.+)\s+Operation: new = (.+)\s+Test: divisible by (\d+)\s+If true: throw to monkey (\d+)\s+If false: throw to monkey (\d+)`)
	opRe     = regexp.MustCompile(`old . (\d+)`)
)

func parse(input string) []*monkey {
	var monkeys []*monkey
	for _, m := range strings.Split(input, "\n\n") {
		caps := monkeyRe.FindStringSubmatch(m)
		m := &monkey{}
		for _, item := range strings.Split(caps[1], ",") {
			num, _ := strconv.Atoi(strings.TrimSpace(item))
			m.items = append(m.items, num)
		}
		m.op = caps[2]
		m.test, _ = strconv.Atoi(caps[3])
		m.ifTrue, _ = strconv.Atoi(caps[4])
		m.ifFalse, _ = strconv.Atoi(caps[5])
		monkeys = append(monkeys, m)
	}
	return monkeys
}

func lcm(a, b int) int {
	// https://rosettacode.org/wiki/Least_common_multiple#Go
	// we know numbers from input are < int32
	m := big.NewInt(int64(a))
	n := big.NewInt(int64(b))
	z := big.Int{}
	return int(z.Mul(z.Div(m, z.GCD(nil, nil, m, n)), n).Int64())
}
