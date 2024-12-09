package main

import (
	_ "embed"
	"fmt"

	"hora.dev/aoc/utils"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func main() {
	part1(input)
}

func part1(input string) {
	ans := checksum(compact(parse(input)))
	fmt.Printf("part 1: %v\n", ans)
}

func checksum(data []int) int64 {
	var result int64
	for i, val := range data {
		if val == -1 {
			break
		}
		result += int64(i) * int64(val)
	}
	return result
}

func compact(data []int) []int {
	freeIdx := nextFreeIdx(data, 0)
	for i := len(data) - 1; i >= 0; i-- {
		if freeIdx >= i {
			break
		}
		data[i], data[freeIdx] = data[freeIdx], data[i]
		freeIdx = nextFreeIdx(data, freeIdx)
	}
	return data
}

func nextFreeIdx(data []int, start int) int {
	for i := start; i < len(data); i++ {
		if data[i] == -1 {
			return i
		}
	}
	panic("no free space?")
}

func parse(input string) []int {
	ints := utils.IntList(input, "")
	var expanded []int
	idx := 0
	for i, val := range ints {
		if i%2 == 0 {
			for range val {
				expanded = append(expanded, idx)
			}
			idx += 1
		} else {
			for range val {
				expanded = append(expanded, -1)
			}
		}
	}
	return expanded
}
