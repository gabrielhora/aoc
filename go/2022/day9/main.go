package main

import (
	_ "embed"
	"fmt"
	"hora.dev/aoc/2022/utils/set"
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
	moves := parse(input)
	fmt.Printf("part 1: %d\n", moveRope(moves, 2))
}

func part2(input string) {
	moves := parse(input)
	fmt.Printf("part 2: %d\n", moveRope(moves, 10))
}

type coord struct {
	x, y int
}

func moveRope(moves []string, numKnots int) int {
	knots := make([]coord, numKnots)
	visited := set.Set[coord]{}
	for _, m := range moves {
		knots = move(m, knots)
		visited.Push(knots[len(knots)-1])
	}
	return len(visited)
}

func move(m string, knots []coord) []coord {
	head, tail := knots[0], knots[1:]
	var k []coord
	switch m {
	case "R":
		k = append(k, coord{head.x + 1, head.y})
		return moveKnots(append(k, tail...))
	case "L":
		k = append(k, coord{head.x - 1, head.y})
		return moveKnots(append(k, tail...))
	case "U":
		k = append(k, coord{head.x, head.y + 1})
		return moveKnots(append(k, tail...))
	case "D":
		k = append(k, coord{head.x, head.y - 1})
		return moveKnots(append(k, tail...))
	default:
		return knots
	}
}

func moveKnots(knots []coord) []coord {
	if len(knots) == 1 {
		return knots
	}

	h, t, rest := knots[0], knots[1], knots[2:]
	dx := h.x - t.x
	dy := h.y - t.y

	moves := map[coord]coord{
		{-2, -2}: {t.x - 1, t.y - 1},
		{-2, -1}: {t.x - 1, t.y - 1},
		{-2, 0}:  {t.x - 1, t.y},
		{-2, 1}:  {t.x - 1, t.y + 1},
		{-2, 2}:  {t.x - 1, t.y + 1},
		{-1, -2}: {t.x - 1, t.y - 1},
		{-1, 2}:  {t.x - 1, t.y + 1},
		{0, -2}:  {t.x, t.y - 1},
		{0, 2}:   {t.x, t.y + 1},
		{1, -2}:  {t.x + 1, t.y - 1},
		{1, 2}:   {t.x + 1, t.y + 1},
		{2, -2}:  {t.x + 1, t.y - 1},
		{2, -1}:  {t.x + 1, t.y - 1},
		{2, 0}:   {t.x + 1, t.y},
		{2, 1}:   {t.x + 1, t.y + 1},
		{2, 2}:   {t.x + 1, t.y + 1},
	}
	tail, ok := moves[coord{dx, dy}]
	if !ok {
		tail = t
	}

	var ret []coord
	ret = append(ret, h)
	return append(ret, moveKnots(append([]coord{tail}, rest...))...)
}

func parse(input string) []string {
	var res []string
	for _, l := range strings.Split(input, "\n") {
		words := strings.Split(l, " ")
		letter := words[0]
		count, _ := strconv.Atoi(words[1])
		for i := 0; i < count; i++ {
			res = append(res, letter)
		}
	}
	return res
}
