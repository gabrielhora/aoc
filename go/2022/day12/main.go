package main

import (
	_ "embed"
	"fmt"
	"hora.dev/aoc/2022/utils/set"
	"hora.dev/aoc/2022/utils/slice"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	st := newState(input)
	findTarget(&st, []nodeHeight{st.initial})
	fmt.Printf("part 1: %d\n", st.moves)
}

func part2(input string) {
	st := newState(input)

	best := math.MaxInt
	for n, h := range st.nodes {
		if h != 'a' {
			continue
		}

		newSt := st
		newSt.visited = set.Set[node]{}
		newSt.moves = 0
		if findTarget(&newSt, []nodeHeight{{n, h}}) {
			if newSt.moves < best {
				best = newSt.moves
			}
		}
	}

	fmt.Printf("part 2: %d\n", best)
}

func findTarget(st *state, visit []nodeHeight) bool {
	for _, item := range visit {
		if item == st.target {
			return true
		}

		// visit neighbor nodes if possible
		for _, d := range []node{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			targetNode := node{item.node.x + d.x, item.node.y + d.y}
			if targetHeight, ok := st.nodes[targetNode]; ok {
				if targetHeight-item.height <= 1 {
					visit = append(visit, nodeHeight{targetNode, targetHeight})
				}
			}
		}
	}

	if len(visit) == 0 {
		return false
	}

	// remove already visited nodes and update state visited nodes
	var dedupVisit []nodeHeight
	for _, n := range visit {
		if _, ok := st.visited[n.node]; !ok {
			dedupVisit = append(dedupVisit, n)
			st.visited.Push(n.node)
		}
	}

	st.moves += 1
	return findTarget(st, dedupVisit)
}

type nodeHeight struct {
	node   node
	height int
}

type node struct {
	x, y int
}

type state struct {
	nodes   map[node]int
	initial nodeHeight
	target  nodeHeight
	moves   int
	visited set.Set[node]
}

func newState(input string) state {
	var letters [][]rune
	for _, l := range strings.Split(input, "\n") {
		letters = append(letters, slice.Runes(l))
	}
	rows := len(letters)
	cols := len(letters[0])

	initial, target, nodes := nodeHeight{}, nodeHeight{}, map[node]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			weight := int(letters[r][c])
			n := node{c, r}
			if weight == 'S' {
				initial = nodeHeight{n, 'a'}
				weight = 'a'
			}
			if weight == 'E' {
				target = nodeHeight{n, 'z'}
				weight = 'z'
			}
			nodes[n] = weight
		}
	}

	return state{
		nodes:   nodes,
		initial: initial,
		target:  target,
		moves:   0,
		visited: set.Set[node]{},
	}
}
