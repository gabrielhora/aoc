package main

import (
	_ "embed"
	"fmt"
	"slices"
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
	history := parse(input)

	result := 0
	for _, h := range history {
		l := layers([][]int{h}, h)
		next := appendNext(l)
		result += next
	}

	fmt.Printf("part 1: %d\n", result)
}

func part2(input string) {
	history := parse(input)

	result := 0
	for _, h := range history {
		slices.Reverse(h)
		l := layers([][]int{h}, h)
		next := appendNext(l)
		result += next
	}

	fmt.Printf("part 2: %d\n", result)
}

func layers(acc [][]int, current []int) [][]int {
	if allZero(current) {
		acc[len(acc)-1] = append(acc[len(acc)-1], 0)
		return acc
	}
	var diffs []int
	for j := 1; j < len(current); j++ {
		diffs = append(diffs, current[j]-current[j-1])
	}
	return layers(append(acc, diffs), diffs)
}

func appendNext(layers [][]int) int {
	for i := len(layers) - 2; i >= 0; i-- {
		fst := layers[i]
		snd := layers[i+1]
		next := fst[len(fst)-1] + snd[len(snd)-1]
		layers[i] = append(fst, next)
	}
	return layers[0][len(layers[0])-1]
}

func parse(input string) [][]int {
	var result [][]int
	for _, line := range strings.Split(input, "\n") {
		var nums []int
		for _, n := range strings.Split(line, " ") {
			nn, _ := strconv.Atoi(n)
			nums = append(nums, nn)
		}
		result = append(result, nums)
	}
	return result
}

func allZero(xs []int) bool {
	for i := 1; i < len(xs); i++ {
		if xs[i] != 0 {
			return false
		}
	}
	return true
}
