package main

import (
	_ "embed"
	"fmt"
	"hora.dev/aoc/2022/utils/set"
	"hora.dev/aoc/2022/utils/slice"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var pairs [][]string
	for _, l := range lines {
		pairs = append(pairs, []string{l[:len(l)/2], l[len(l)/2:]})
	}

	var inter []rune
	for _, p := range pairs {
		a, b := set.Runes(p[0]), set.Runes(p[1])
		inter = append(inter, a.Intersection(b).Slice()...)
	}

	fmt.Printf("part 1: %d\n", priority(inter))
}

func part2(lines []string) {
	var inter []rune
	for _, g := range slice.Split(lines, 3) {
		a, b, c := set.Runes(g[0]), set.Runes(g[1]), set.Runes(g[2])
		inter = append(inter, a.Intersection(b).Intersection(c).Slice()...)
	}

	fmt.Printf("part 2: %d\n", priority(inter))
}

func priority(letters []rune) int {
	ans := 0
	for _, i := range letters {
		if i >= 97 && i <= 122 {
			ans += int(i) - 96
		} else {
			ans += int(i) - 38
		}
	}
	return ans
}
