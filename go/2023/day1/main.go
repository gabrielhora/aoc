package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	data := parse(input)
	sum := 0
	for _, line := range data {
		ns := findFirstNum(line, false) + findLastNum(line, false)
		n, _ := strconv.Atoi(ns)
		sum += n
	}
	fmt.Printf("part 1: %v\n", sum)
}

func part2(input string) {
	data := parse(input)
	sum := 0
	for _, line := range data {
		ns := findFirstNum(line, true) + findLastNum(line, true)
		n, _ := strconv.Atoi(ns)
		sum += n
	}
	fmt.Printf("part 2: %v\n", sum)
}

func findFirstNum(line []rune, findWords bool) string {
	acc := ""
	for i := 0; i < len(line); i++ {
		if findWords {
			acc += string(line[i])
			if val := hasNumberWord(acc); val != "" {
				return val
			}
		}
		if unicode.IsDigit(line[i]) {
			return string(line[i])
		}
	}
	panic("should have found first number")
}

func findLastNum(line []rune, findWords bool) string {
	acc := ""
	for i := len(line) - 1; i >= 0; i-- {
		if findWords {
			acc = string(line[i]) + acc // prepend
			if val := hasNumberWord(acc); val != "" {
				return val
			}
		}
		if unicode.IsDigit(line[i]) {
			return string(line[i])
		}
	}
	panic("should have found last number")
}

var numWords = map[string]string{"zero": "0", "one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

func hasNumberWord(val string) string {
	for name, num := range numWords {
		if strings.Contains(val, name) {
			return num
		}
	}
	return ""
}

func parse(input string) [][]rune {
	var result [][]rune
	for _, l := range strings.Split(input, "\n") {
		result = append(result, []rune(l))
	}
	return result
}
