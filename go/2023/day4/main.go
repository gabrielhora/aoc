package main

import (
	_ "embed"
	"fmt"
	"golang.org/x/exp/maps"
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
	points := 0
	for _, c := range parse(input) {
		points += c.points()
	}
	fmt.Printf("part 1: %d\n", points)
}

func part2(input string) {
	originalCards := parse(input)

	result := map[int][]card{}
	for n, c := range originalCards {
		result[n] = append(result[n], c)
	}

	// duplicate cards
	for i := 1; i < len(result)+1; i++ {
		numDups := len(originalCards[i].matches())
		for range result[i] {
			for j := i + 1; j <= i+numDups; j++ {
				result[j] = append(result[j], originalCards[j])
			}
		}
	}

	count := 0
	for _, cs := range result {
		count += len(cs)
	}

	fmt.Printf("part 2: %d\n", count)
}

type card struct {
	number  int
	numbers map[int]struct{}
	winning map[int]struct{}
}

func (c card) matches() []int {
	var matches []int
	for _, n := range maps.Keys(c.winning) {
		if _, ok := c.numbers[n]; ok {
			matches = append(matches, n)
		}
	}
	return matches
}

func (c card) points() int {
	m := c.matches()
	return int(math.Pow(2, float64(len(m)-1)))
}

var space = regexp.MustCompile(`\s+`)

func parse(input string) map[int]card {
	cards := map[int]card{}
	for i, line := range strings.Split(input, "\n") {
		line = line[strings.Index(line, ":")+1:]
		parts := strings.Split(line, "|")
		c := card{
			numbers: map[int]struct{}{},
			winning: map[int]struct{}{},
		}

		for _, ns := range space.Split(strings.TrimSpace(parts[0]), -1) {
			n, _ := strconv.Atoi(ns)
			c.numbers[n] = struct{}{}
		}

		for _, ns := range space.Split(strings.TrimSpace(parts[1]), -1) {
			n, _ := strconv.Atoi(ns)
			c.winning[n] = struct{}{}
		}

		cards[i+1] = c
	}
	return cards
}
