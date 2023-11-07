package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func intList(listOfNums string) []int {
	lines := strings.Split(listOfNums, "\n")
	var nums []int
	for _, s := range lines {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}
	return nums
}

func part1(input string) int {
	nums := intList(input)
	for _, n1 := range nums {
		for _, n2 := range nums {
			if n1+n2 == 2020 {
				return n1 * n2
			}
		}
	}
	return 0
}

func part2(input string) int {
	nums := intList(input)
	for _, n1 := range nums {
		for _, n2 := range nums {
			for _, n3 := range nums {
				if n1+n2+n3 == 2020 {
					return n1 * n2 * n3
				}
			}
		}
	}
	return 0
}

func main() {
	println(part1(input))
	println(part2(input))
}
