package main

import (
	_ "embed"
	"fmt"
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
	result := 0
	for _, g := range parse(input) {
		if g.isValid() {
			result += g.id
		}
	}
	fmt.Printf("part 1: %d\n", result)
}

func part2(input string) {
	result := 0
	for _, g := range parse(input) {
		result += g.power()
	}
	fmt.Printf("part 2: %d\n", result)
}

type game struct {
	id    int
	plays []map[string]int
}

func (g game) isValid() bool {
	for _, play := range g.plays {
		red, _ := play["red"]
		green, _ := play["green"]
		blue, _ := play["blue"]
		if red > 12 || green > 13 || blue > 14 {
			return false
		}
	}
	return true
}

func (g game) power() int {
	mr := 0
	mg := 0
	mb := 0
	for _, play := range g.plays {
		red, _ := play["red"]
		if red > mr {
			mr = red
		}
		green, _ := play["green"]
		if green > mg {
			mg = green
		}
		blue, _ := play["blue"]
		if blue > mb {
			mb = blue
		}
	}
	return mr * mg * mb
}

func parse(input string) []game {
	lines := strings.Split(input, "\n")
	games := make([]game, 0, len(lines))
	for _, line := range lines {
		comma := strings.Index(line, ":")
		id, _ := strconv.Atoi(line[5:comma])
		plays := strings.Split(line[comma+1:], ";")
		g := game{id: id}
		for _, p := range plays {
			colorsPlayed := map[string]int{}
			for _, c := range strings.Split(p, ",") {
				cp := strings.Split(strings.TrimSpace(c), " ")
				n, _ := strconv.Atoi(cp[0])
				colorsPlayed[cp[1]] = n
			}
			g.plays = append(g.plays, colorsPlayed)
		}
		games = append(games, g)
	}
	return games
}
