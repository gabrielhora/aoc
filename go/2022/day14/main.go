package main

import (
	_ "embed"
	"fmt"
	"hora.dev/aoc/2022/utils/set"
	"hora.dev/aoc/2022/utils/slice"
	"math"
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
	blocked := parse(input)
	// find the void limit
	yLim := 0
	for b := range blocked {
		if b[1] > yLim {
			yLim = b[1]
		}
	}
	origin := [2]int{500, 0}

	// drop sand until one of them fall to the void
	ans := 0
	for drop(blocked, origin, yLim) {
		ans += 1
	}
	fmt.Printf("part 1: %d\n", ans)
}

func part2(input string) {
	blocked := parse(input)
	// find the floor
	yLim := 0
	for b := range blocked {
		if b[1] > yLim {
			yLim = b[1]
		}
	}
	yLim += 2
	origin := [2]int{500, 0}

	// drop sand until origin is blocked
	ans := 0
	for !blocked.Contains(origin) {
		drop(blocked, origin, yLim)
		ans += 1
	}
	fmt.Printf("part 2: %d\n", ans)
}

// drop returns false if sand goes over the limit
func drop(blocked set.Set[[2]int], origin [2]int, lim int) bool {
	down := [2]int{origin[0], origin[1] + 1}
	left := [2]int{origin[0] - 1, origin[1] + 1}
	right := [2]int{origin[0] + 1, origin[1] + 1}

	if origin[1] >= lim {
		blocked.Push([2]int{origin[0], origin[1] - 1})
		return false
	} else if !blocked.Contains(down) {
		return drop(blocked, down, lim)
	} else if !blocked.Contains(left) {
		return drop(blocked, left, lim)
	} else if !blocked.Contains(right) {
		return drop(blocked, right, lim)
	} else {
		blocked.Push(origin)
		return true
	}
}

func parse(input string) set.Set[[2]int] {
	var res [][2]int

	for _, line := range strings.Split(input, "\n") {
		var pairs [][2]int
		for _, coord := range strings.Split(line, " -> ") {
			p := strings.Split(coord, ",")
			x, _ := strconv.Atoi(p[0])
			y, _ := strconv.Atoi(p[1])
			pairs = append(pairs, [2]int{x, y})
		}

		for _, ps := range slice.SlidingWindow(pairs, 2) {
			x1 := min(ps[0][0], ps[1][0])
			x2 := max(ps[0][0], ps[1][0])
			y1 := min(ps[0][1], ps[1][1])
			y2 := max(ps[0][1], ps[1][1])

			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					res = append(res, [2]int{x, y})
				}
			}
		}
	}

	return set.FromValues(res...)
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
