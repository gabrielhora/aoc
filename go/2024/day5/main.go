package main

import (
	_ "embed"
	"fmt"
	"hora.dev/aoc/utils"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

//go:embed sample.txt
var sample string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	parts := strings.Split(input, "\n\n")
	orders := parseOrders(parts[0])
	updates := parseUpdates(parts[1])
	pageDeps := pageDependencies(orders)
	valids, _ := splitValidAndInvalid(updates, pageDeps)

	var ans int64
	for _, upd := range valids {
		ans += upd[(len(upd) / 2)]
	}
	fmt.Printf("part 1: %d\n", ans)
}

func part2(input string) {
	parts := strings.Split(input, "\n\n")
	orders := parseOrders(parts[0])
	updates := parseUpdates(parts[1])
	pageDeps := pageDependencies(orders)
	_, invalids := splitValidAndInvalid(updates, pageDeps)

	var fixed [][]int64
	for _, pages := range invalids {
		fixed = append(fixed, fixOrder(pages, pageDeps))
	}

	var ans int64
	for _, pages := range fixed {
		ans += pages[(len(pages) / 2)]
	}
	fmt.Printf("part 2: %d\n", ans)
}

// fixOrder loops and swaps invalid items until in the correct order
func fixOrder(pages []int64, deps map[int64][]int64) []int64 {
	startIdx := 0
	for {
		isValid := true
		for i := startIdx; i < len(pages); i++ {
			invalidIdx, valid := isInCorrectOrder(pages, i, deps)
			if valid {
				continue
			}
			isValid = false
			pages[i], pages[invalidIdx] = pages[invalidIdx], pages[i]
			startIdx = i
			break
		}
		if isValid {
			break
		}
	}
	return pages
}

func splitValidAndInvalid(updates [][]int64, deps map[int64][]int64) (valid [][]int64, invalid [][]int64) {
	for _, pages := range updates {
		isValid := true
		for pageIdx := range pages {
			if _, valid := isInCorrectOrder(pages, pageIdx, deps); !valid {
				isValid = false
				break
			}
		}
		if isValid {
			valid = append(valid, pages)
		} else {
			invalid = append(invalid, pages)
		}
	}
	return
}

func isInCorrectOrder(pages []int64, pageIdx int, deps map[int64][]int64) (int, bool) {
	pageDeps := deps[pages[pageIdx]]
	for idx, nextPage := range pages[pageIdx+1:] {
		if !slices.Contains(pageDeps, nextPage) {
			return pageIdx + idx + 1, false
		}
	}
	return -1, true
}

// pageDependencies builds a dependency list, as in pages that must come first, for each page
func pageDependencies(orders [][]int64) map[int64][]int64 {
	deps := make(map[int64][]int64)
	for _, ord := range orders {
		if _, ok := deps[ord[0]]; ok {
			deps[ord[0]] = append(deps[ord[0]], ord[1])
		} else {
			deps[ord[0]] = []int64{ord[1]}
		}
	}
	return deps
}

func parseOrders(input string) [][]int64 {
	var orders [][]int64
	for _, pair := range strings.Split(input, "\n") {
		orders = append(orders, utils.IntList(pair, "|"))
	}
	return orders
}

func parseUpdates(input string) [][]int64 {
	var updates [][]int64
	for _, csv := range strings.Split(input, "\n") {
		updates = append(updates, utils.IntList(csv, ","))
	}
	return updates
}
