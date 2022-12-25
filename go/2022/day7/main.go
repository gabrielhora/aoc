package main

import (
	_ "embed"
	"fmt"
	"path"
	"regexp"
	"sort"
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
	s := state{cwd: "/", tree: map[string]int{}}
	for _, cmd := range strings.Split(input, "\n") {
		s.run(cmd)
	}
	ans := 0
	for _, size := range s.tree {
		if size <= 100000 {
			ans += size
		}
	}
	fmt.Printf("part 1: %d\n", ans)
}

func part2(input string) {
	s := state{cwd: "/", tree: map[string]int{}}
	for _, cmd := range strings.Split(input, "\n") {
		s.run(cmd)
	}

	unused := 70000000 - s.tree["/"]
	var sizes []int
	for _, s := range s.tree {
		sizes = append(sizes, s)
	}
	sort.Ints(sizes)
	ans := 0
	for _, s := range sizes {
		if unused+s >= 30000000 {
			ans = s
			break
		}
	}
	fmt.Printf("part 2: %d\n", ans)
}

var (
	cdIn     = regexp.MustCompile(`\$ cd (.*)`)
	cdOut    = regexp.MustCompile(`\$ cd \.\.`)
	fileLine = regexp.MustCompile(`(\d+) (.*)`)
)

type state struct {
	cwd  string
	tree map[string]int
}

func (s *state) run(cmd string) {
	switch {
	case cdIn.MatchString(cmd):
		caps := cdIn.FindStringSubmatch(cmd)
		s.cwd = path.Join(s.cwd, caps[1])
	case cdOut.MatchString(cmd):
		s.cwd = path.Dir(s.cwd)
	case fileLine.MatchString(cmd):
		caps := fileLine.FindStringSubmatch(cmd)
		size, _ := strconv.Atoi(caps[1])
		s.addSize(size, s.cwd)
	}
}

func (s *state) addSize(size int, dir string) {
	s.tree[dir] += size
	if dir == "/" {
		return
	}
	s.addSize(size, path.Dir(dir))
}
