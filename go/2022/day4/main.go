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
	data := parse(input)
	part1(data)
	part2(data)
}

func part1(data [][]int) {
	ans := 0
	for _, d := range data {
		a := set.FromRange(d[0], d[1])
		b := set.FromRange(d[2], d[3])
		if a.Intersection(b).Equal(b) || b.Intersection(a).Equal(a) {
			ans += 1
		}
	}
	fmt.Printf("part 1: %d\n", ans)
}

func part2(data [][]int) {
	ans := 0
	for _, d := range data {
		a := set.FromRange(d[0], d[1])
		b := set.FromRange(d[2], d[3])
		if len(a.Intersection(b)) != 0 {
			ans += 1
		}
	}
	fmt.Printf("part 2: %d\n", ans)
}

func parse(input string) [][]int {
	var res [][]int
	for _, line := range strings.Split(input, "\n") {
		p := strings.Split(line, ",")
		p0 := strings.Split(p[0], "-")
		p1 := strings.Split(p[1], "-")

		a, _ := strconv.Atoi(p0[0])
		b, _ := strconv.Atoi(p0[1])
		c, _ := strconv.Atoi(p1[0])
		d, _ := strconv.Atoi(p1[1])

		res = append(res, []int{a, b, c, d})
	}
	return res
}
