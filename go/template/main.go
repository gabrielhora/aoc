package main

import _ "embed"

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func main() {
	part1(example)
	part2(example)
}

func part1(input string) {
	// TODO: implement me
}

func part2(input string) {
	// TODO: implement me
}
