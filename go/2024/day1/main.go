package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed day1.txt
var day1 string

func main() {
	part1()
	part2()
}

func part1() {
	list1, list2 := splitLists(day1)
	slices.Sort(list1)
	slices.Sort(list2)

	var result float64
	for i := 0; i < len(list1); i++ {
		num1 := float64(list1[i])
		num2 := float64(list2[i])
		result += math.Abs(num1 - num2)
	}

	fmt.Printf("part 1: %v\n", int64(result))
}

func part2() {
	list1, list2 := splitLists(day1)

	// zip list2 with counts
	map2 := make(map[int64]int64)
	for i := 0; i < len(list2); i++ {
		if _, ok := map2[list2[i]]; ok {
			map2[list2[i]] += 1
		} else {
			map2[list2[i]] = 1
		}
	}

	var similarity int64
	for i := 0; i < len(list1); i++ {
		if count, ok := map2[list1[i]]; ok {
			similarity += list1[i] * count
		}
	}

	fmt.Printf("part 2: %v\n", similarity)
}

var spaces = regexp.MustCompile("\\s+")

func splitLists(input string) ([]int64, []int64) {
	var list1, list2 []int64
	for _, line := range strings.Split(input, "\n") {
		nums := spaces.Split(line, -1)
		if len(nums) == 0 {
			continue
		}
		num1, _ := strconv.ParseInt(nums[0], 10, 64)
		num2, _ := strconv.ParseInt(nums[1], 10, 64)
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	return list1, list2
}
