package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"hora.dev/aoc/2022/utils/set"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	cubes := parse(input)
	ans := 0
	for c := range cubes {
		if !cubes.Contains(cube{c.x + 1, c.y, c.z}) {
			ans += 1
		}
		if !cubes.Contains(cube{c.x - 1, c.y, c.z}) {
			ans += 1
		}
		if !cubes.Contains(cube{c.x, c.y + 1, c.z}) {
			ans += 1
		}
		if !cubes.Contains(cube{c.x, c.y - 1, c.z}) {
			ans += 1
		}
		if !cubes.Contains(cube{c.x, c.y, c.z + 1}) {
			ans += 1
		}
		if !cubes.Contains(cube{c.x, c.y, c.z - 1}) {
			ans += 1
		}
	}
	fmt.Printf("part 1: %d\n", ans)
}

func part2(input string) {
	cs := parse(input)
	ans := 0
	for c := range cs {
		if escape(cube{c.x + 1, c.y, c.z}, cs) {
			ans += 1
		}
		if escape(cube{c.x - 1, c.y, c.z}, cs) {
			ans += 1
		}
		if escape(cube{c.x, c.y + 1, c.z}, cs) {
			ans += 1
		}
		if escape(cube{c.x, c.y - 1, c.z}, cs) {
			ans += 1
		}
		if escape(cube{c.x, c.y, c.z + 1}, cs) {
			ans += 1
		}
		if escape(cube{c.x, c.y, c.z - 1}, cs) {
			ans += 1
		}
	}
	fmt.Printf("part 2: %d\n", ans)
}

func escape(c cube, cs set.Set[cube]) bool {
	visited := set.Set[cube]{}
	q := []cube{c}

	var cur cube
	for len(q) > 0 {
		cur, q = q[0], q[1:]
		if cs.Contains(cur) || visited.Contains(cur) {
			continue
		}
		visited.Push(cur)
		if len(visited) > 2000 {
			return true
		}

		q = append(q, cube{cur.x + 1, cur.y, cur.z})
		q = append(q, cube{cur.x - 1, cur.y, cur.z})
		q = append(q, cube{cur.x, cur.y + 1, cur.z})
		q = append(q, cube{cur.x, cur.y - 1, cur.z})
		q = append(q, cube{cur.x, cur.y, cur.z + 1})
		q = append(q, cube{cur.x, cur.y, cur.z - 1})
	}

	return false
}

type cube struct {
	x, y, z int
}

func parse(input string) set.Set[cube] {
	cubes := set.Set[cube]{}
	for _, line := range strings.Split(input, "\n") {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		cubes.Push(cube{x, y, z})
	}
	return cubes
}
