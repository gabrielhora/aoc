package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	bags := parse(input)
	result := 0
	for _, b := range bags {
		if findBag(bags, b.name, "shiny gold") {
			result += 1
		}
	}
	fmt.Printf("part 1: %d\n", result)
}

func part2(input string) {
	result := countBags(parse(input), "shiny gold") - 1 // -1 for the initial bag
	fmt.Printf("part 2: %d", result)
}

func findBag(bags map[string]bag, start string, lookup string) bool {
	for _, b := range bags[start].contents {
		if b.name == lookup || findBag(bags, b.name, lookup) {
			return true
		}
	}
	return false
}

func countBags(bags map[string]bag, lookup string) int {
	acc := 1
	for _, b := range bags[lookup].contents {
		acc += b.count * countBags(bags, b.name)
	}
	return acc
}

type bag struct {
	name     string
	count    int
	contents []bag
}

func parse(input string) map[string]bag {
	lines := strings.Split(input, "\n")
	bags := map[string]bag{}
	for _, line := range lines {
		b := parseBag(line)
		bags[b.name] = b
	}
	return bags
}

func parseBag(line string) bag {
	var bagRe = regexp.MustCompile(`^\s*(\w+ \w+)|(\d+)\s*(\w+ \w+)`)
	matches := bagRe.FindAllStringSubmatch(line, -1)
	b := bag{name: matches[0][0]}
	for _, match := range matches[1:] {
		num, _ := strconv.Atoi(match[2])
		b.contents = append(b.contents, bag{
			name:  match[3],
			count: num,
		})
	}
	return b
}
