package main

import (
	_ "embed"
	"errors"
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
	cmds := parse(input)
	acc, _ := execute(0, []int{}, 0, cmds)
	fmt.Printf("part 1: %d\n", acc)
}

func part2(input string) {
	cmds := parse(input)
	acc := fixLoopWithBruteForce(cmds)
	fmt.Printf("part 2: %d\n", acc)
}

var errLoop = errors.New("looping")

func execute(acc int, visited []int, index int, cmds []command) (int, error) {
	if index == len(cmds) {
		return acc, nil
	}

	// are we looping?
	for _, i := range visited {
		if i == index {
			return acc, errLoop
		}
	}

	cmd := cmds[index]
	visited = append(visited, index)
	switch cmd.op {
	case "nop":
		return execute(acc, visited, index+1, cmds)
	case "acc":
		acc += cmd.num
		return execute(acc, visited, index+1, cmds)
	case "jmp":
		return execute(acc, visited, index+cmd.num, cmds)
	default:
		panic("invalid instruction")
	}
}

func fixLoopWithBruteForce(cmds []command) int {
	changedCmds, i := changeNextNopOrJmp(cmds, 0)
	for {
		acc, err := execute(0, []int{}, 0, changedCmds)
		if err == nil {
			return acc
		}
		changedCmds, i = changeNextNopOrJmp(cmds, i+1)
	}
}

func changeNextNopOrJmp(cmds []command, i int) ([]command, int) {
	newCmds := append([]command{}, cmds...)
	for j := i; j < len(newCmds); j++ {
		cmd := &newCmds[j]
		if cmd.op == "nop" {
			cmd.op = "jmp"
			return newCmds, j
		}
		if cmd.op == "jmp" {
			cmd.op = "nop"
			return newCmds, j
		}
	}
	return newCmds, -1
}

type command struct {
	op  string
	num int
}

func parse(input string) []command {
	cmdRe := regexp.MustCompile(`(\w+)\s([+-]\d+)`)
	lines := strings.Split(input, "\n")
	var cmds []command
	for _, line := range lines {
		matches := cmdRe.FindAllStringSubmatch(line, -1)
		num, _ := strconv.Atoi(matches[0][2])
		cmds = append(cmds, command{
			op:  matches[0][1],
			num: num,
		})
	}
	return cmds
}
