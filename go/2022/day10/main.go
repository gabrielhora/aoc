package main

import (
	_ "embed"
	"fmt"
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
	cycles := execute(input)
	ans := 20*cycles[19] +
		60*cycles[59] +
		100*cycles[99] +
		140*cycles[139] +
		180*cycles[179] +
		220*cycles[219]
	fmt.Printf("part 1: %d\n", ans)
}

func part2(input string) {
	cycles := execute(input)

	fmt.Println("part 2:")
	var pixels []string
	for cycle, x := range cycles {
		xt := (cycle/40)*40 + x
		if abs(xt-cycle) <= 1 {
			pixels = append(pixels, "#")
		} else {
			pixels = append(pixels, " ")
		}
	}
	for _, l := range slice.Split(pixels, 40) {
		fmt.Println(strings.Join(l, ""))
	}
}

func execute(cmds string) []int {
	x := 1
	xs := []int{1}
	for _, cmd := range strings.Split(cmds, "\n") {
		if strings.HasPrefix(cmd, "addx") {
			num, _ := strconv.Atoi(cmd[5:])
			// value only change after the second cycle
			xs = append(xs, []int{x, x + num}...)
			x += num
		} else {
			xs = append(xs, x)
		}
	}
	return xs
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
