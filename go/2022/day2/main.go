package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	data := parse(input)
	part1(data)
	part2(data)
}

func parse(input string) [][]string {
	var res [][]string
	for _, l := range strings.Split(input, "\n") {
		res = append(res, strings.Split(l, " "))
	}
	return res
}

func part1(data [][]string) {
	ans := 0
	for _, pair := range data {
		a, b := pair[0], pair[1]
		switch {
		case a == "A" && b == "X":
			ans += 1 + 3
		case a == "A" && b == "Y":
			ans += 2 + 6
		case a == "A" && b == "Z":
			ans += 3 + 0
		case a == "B" && b == "X":
			ans += 1 + 0
		case a == "B" && b == "Y":
			ans += 2 + 3
		case a == "B" && b == "Z":
			ans += 3 + 6
		case a == "C" && b == "X":
			ans += 1 + 6
		case a == "C" && b == "Y":
			ans += 2 + 0
		case a == "C" && b == "Z":
			ans += 3 + 3
		}
	}
	fmt.Printf("part 1: %d\n", ans)
}

func part2(data [][]string) {
	ans := 0
	for _, pair := range data {
		a, b := pair[0], pair[1]
		switch {
		case a == "A" && b == "X":
			ans += 3 + 0
		case a == "A" && b == "Y":
			ans += 1 + 3
		case a == "A" && b == "Z":
			ans += 2 + 6
		case a == "B" && b == "X":
			ans += 1 + 0
		case a == "B" && b == "Y":
			ans += 2 + 3
		case a == "B" && b == "Z":
			ans += 3 + 6
		case a == "C" && b == "X":
			ans += 2 + 0
		case a == "C" && b == "Y":
			ans += 3 + 3
		case a == "C" && b == "Z":
			ans += 1 + 6
		}
	}
	fmt.Printf("part 2: %d\n", ans)
}
