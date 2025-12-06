package main

import (
	_ "embed"
	"fmt"
	"strings"

	"hora.dev/aoc/2022/utils/set"
	"hora.dev/aoc/utils"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	ranges, ingredients := parse(input)
	fresh := 0
	for _, i := range ingredients {
		for _, r := range ranges {
			if i >= r[0] && i <= r[1] {
				fresh++
				break
			}
		}
	}
	fmt.Printf("Part 1: %v\n", fresh)
}

func part2(input string) {
	ranges, _ := parse(input)
	ranges = merge(ranges)
	count := int64(0)
	for _, r := range ranges {
		count += r[1] - r[0] + 1
	}
	fmt.Printf("Part 2: %v\n", count)
}

func parse(input string) ([][2]int64, []int64) {
	parts := strings.Split(input, "\n\n")
	var ranges [][2]int64
	for l := range strings.Lines(parts[0]) {
		p := utils.IntList[int64](l, "-")
		ranges = append(ranges, [2]int64{p[0], p[1]})
	}
	ingredients := utils.IntList[int64](parts[1], "\n")
	return ranges, ingredients
}

func merge(ranges [][2]int64) [][2]int64 {
	mset := set.Set[[2]int64]{}
	for _, r := range ranges {
		mset.Push(r)
	}

	for {
		mc := 0
		for r1 := range mset {
			s1 := r1[0]
			e1 := r1[1]
			origk := [2]int64{s1, e1}
			m := false

			for r2 := range mset {
				s2 := r2[0]
				e2 := r2[1]
				if s1 == s2 && e1 == e2 {
					continue
				}
				// r2 starts after r1
				if s2 >= s1 && s2 <= e1 {
					m = true // no change
				}
				// r2 ends before r1
				if e2 >= s1 && e2 <= e1 {
					m = true
				}
				// r2 starts before r1 and ends inside r1
				if s2 < s1 && e2 >= s1 && e2 <= e1 {
					m = true
					s1 = s2 // change start to smaller
				}
				// r2 starts inside r1 and ends after r1
				if s2 > s1 && s2 <= e1 && e2 > e1 {
					m = true
					e1 = e2 // change end to bigger
				}
				// r2 around r1
				if s2 <= s1 && e2 >= e1 {
					m = true
					s1 = s2
					e1 = e2
				}
			}

			if m {
				newk := [2]int64{s1, e1}
				mset.Del(origk)
				mset.Push(newk)
				mc++
			}
		}
		if mc == 0 {
			break
		}
	}

	return mset.Slice()
}
