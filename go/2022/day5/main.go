package main

import (
	_ "embed"
	"fmt"
	"math"
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
	parts := strings.Split(input, "\n\n")
	stacks := parseStacks(parts[0])
	moves := parseMoves(parts[1])

	// expand moves so always moves one crate at a time
	var expandedMoves []move
	for _, m := range moves {
		for i := 0; i < m.qty; i++ {
			expandedMoves = append(expandedMoves, move{m.from, m.to, 1})
		}
	}

	stacks = moveCrates(expandedMoves, stacks)
	ans := ""
	for _, s := range stacks {
		ans += s[0]
	}
	fmt.Printf("part 1: %s\n", ans)
}

func part2(input string) {
	parts := strings.Split(input, "\n\n")
	stacks := parseStacks(parts[0])
	moves := parseMoves(parts[1])

	stacks = moveCrates(moves, stacks)
	ans := ""
	for _, s := range stacks {
		ans += s[0]
	}
	fmt.Printf("part 2: %s\n", ans)
}

func moveCrates(moves []move, stacks [][]string) [][]string {
	for _, m := range moves {
		head, tail := stacks[m.from][:m.qty], stacks[m.from][m.qty:]
		stacks[m.from] = tail

		var to []string
		to = append(to, head...)
		to = append(to, stacks[m.to]...)
		stacks[m.to] = to
	}
	return stacks
}

var moveRe = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

type move struct {
	from, to, qty int
}

func parseMoves(input string) []move {
	var moves []move
	for _, l := range strings.Split(input, "\n") {
		nums := moveRe.FindStringSubmatch(l)
		qty, _ := strconv.Atoi(nums[1])
		from, _ := strconv.Atoi(nums[2])
		to, _ := strconv.Atoi(nums[3])
		moves = append(moves, move{from - 1, to - 1, qty})
	}
	return moves
}

func parseStacks(input string) [][]string {
	cols, _ := strconv.Atoi(string(input[len(input)-1]))
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	rows := len(lines)
	replacer := strings.NewReplacer("[", "", "]", "", " ", "")

	// read letters on a column by column basis
	var stacks [][]string
	for col := 0; col < cols; col++ {
		var stack []string
		for row := 0; row < rows; row++ {
			line := lines[row]
			letter := replacer.Replace(substr(line, col*4, 4))
			if letter != "" {
				stack = append(stack, letter)
			}
		}
		stacks = append(stacks, stack)
	}

	return stacks
}

func substr(str string, start, max int) string {
	if start > len(str) {
		return ""
	}
	ss := str[start:]
	max = int(math.Min(float64(max), float64(len(ss))))
	return ss[:max]
}
