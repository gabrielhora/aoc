package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

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
	ranges := parse(input)
	var res int64
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			if isInvalidPart1(i) {
				res += i
			}
		}
	}
	fmt.Printf("part 1: %v\n", res)
}

func part2(input string) {
	ranges := parse(input)
	var res int64
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			if isInvalidPart2(i) {
				res += i
			}
		}
	}
	fmt.Printf("part 2: %v\n", res)
}

func isInvalidPart1(val int64) bool {
	cnt := utils.CountDigits(val)
	if cnt < 2 || cnt%2 != 0 {
		return false
	}
	cs := utils.ChunkNumber(val, cnt, cnt/2)
	return cs[0] == cs[1]
}

func isInvalidPart2(val int64) bool {
	cnt := utils.CountDigits(val)
	if cnt == 1 {
		return false
	}
	for i := int64(1); i < cnt; i++ {
		cmap := map[int64]*struct{}{}
		for _, c := range utils.ChunkNumber(val, cnt, i) {
			cmap[c] = nil
		}
		if len(cmap) == 1 { // all chunks are the number
			return true
		}
	}
	return false
}

func parse(input string) [][2]int64 {
	var ranges [][2]int64
	for r := range strings.SplitSeq(strings.TrimSpace(input), ",") {
		rr := strings.Split(r, "-")
		r1, _ := strconv.ParseInt(rr[0], 10, 64)
		r2, _ := strconv.ParseInt(rr[1], 10, 64)
		ranges = append(ranges, [2]int64{min(r1, r2), max(r1, r2)})
	}
	return ranges
}
