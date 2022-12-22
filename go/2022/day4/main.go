package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

type parsedInput struct {
	elf1Start int
	elf1End   int
	elf2Start int
	elf2End   int
}

func parse(input string) []parsedInput {
	lines := strings.Split(input, "\n")
	var result []parsedInput
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		p1Start, _ := strconv.Atoi(strings.Split(pairs[0], "-")[0])
		p1End, _ := strconv.Atoi(strings.Split(pairs[0], "-")[1])
		p2Start, _ := strconv.Atoi(strings.Split(pairs[1], "-")[0])
		p2End, _ := strconv.Atoi(strings.Split(pairs[1], "-")[1])

		result = append(result, parsedInput{
			elf1Start: p1Start,
			elf1End:   p1End,
			elf2Start: p2Start,
			elf2End:   p2End,
		})
	}
	return result
}

func main() {
	fmt.Printf("%+v", parse(example))
}
