package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"hora.dev/aoc/2022/utils/math"
	"hora.dev/aoc/2022/utils/slice"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	nums := parse(input)
	ans := solve(mix(nums, nums))
	fmt.Printf("part 1: %d\n", ans)
}

func part2(input string) {
	nums := parse(input)
	for i := range nums {
		nums[i][0] = nums[i][0] * 811589153
	}

	mixed := append([][2]int{}, nums...)
	for i := 0; i < 10; i++ {
		mixed = mix(nums, mixed)
	}

	fmt.Printf("part 2: %d\n", solve(mixed))
}

func solve(data [][2]int) int {
	size := len(data)
	zero := 0
	for i, d := range data {
		if d[0] == 0 {
			zero = i
			break
		}
	}
	a := data[math.Mod(zero+1000, size)][0]
	b := data[math.Mod(zero+2000, size)][0]
	c := data[math.Mod(zero+3000, size)][0]
	return a + b + c
}

func mix(original, mixed [][2]int) [][2]int {
	size := len(original)
	res := append([][2]int{}, mixed...)
	for _, d := range original {
		val := d[0]
		curIdx := slice.FindIndex(res, d)
		newIdx := math.Mod(curIdx+val, size-1)
		if newIdx == 0 {
			newIdx = size - 1
		}
		res = append(res[:curIdx], res[curIdx+1:]...)
		res = append(res[:newIdx], append([][2]int{d}, res[newIdx:]...)...)
	}
	return res
}

func parse(input string) [][2]int {
	var res [][2]int
	for i, l := range strings.Split(input, "\n") {
		n, _ := strconv.Atoi(l)
		res = append(res, [2]int{n, i})
	}
	return res
}
