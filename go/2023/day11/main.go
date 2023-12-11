package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	image := parse(input)
	galaxies := findGalaxies(image)
	dist := sumDistances(image, galaxies, 1)
	fmt.Printf("part 1: %d\n", dist)
}

func part2(input string) {
	image := parse(input)
	galaxies := findGalaxies(image)
	dist := sumDistances(image, galaxies, 1000000)
	fmt.Printf("part 2: %d\n", dist)
}

func sumDistances(image [][]string, galaxies []point, weight int) int {
	weight = max(1, weight-1)
	totalDistance := 0
	p := pairs(galaxies)
	for i, pair := range p {
		if i%1000 == 0 {
			fmt.Printf("%d / %d\n", i, len(p))
		}
		nr := newRowsBetween(image, pair[0].y, pair[1].y)
		nc := newColsBetween(image, pair[0].x, pair[1].x)
		dist := shortestPath(image, pair[0], pair[1])
		totalDistance += dist + nr*weight + nc*weight
	}
	return totalDistance
}

func newRowsBetween(image [][]string, p1, p2 int) int {
	count := 0
	s, e := min(p1, p2), max(p1, p2)
	for y := s; y < e; y++ {
		if slices.Index(image[y], "#") < 0 {
			count++
		}
	}
	return count
}

func newColsBetween(image [][]string, p1, p2 int) int {
	count := 0
	s, e := min(p1, p2), max(p1, p2)
	for x := s; x < e; x++ {
		duplicate := true
		for _, row := range image {
			if row[x] != "." {
				duplicate = false
				break
			}
		}
		if duplicate {
			count++
		}
	}
	return count
}

type point struct{ x, y int }

func findGalaxies(image [][]string) []point {
	var galaxies []point
	for y := 0; y < len(image); y++ {
		for x := 0; x < len(image[y]); x++ {
			if image[y][x] == "#" {
				galaxies = append(galaxies, point{x, y})
			}
		}
	}
	return galaxies
}

func pairs(points []point) [][2]point {
	var result [][2]point
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			result = append(result, [2]point{points[i], points[j]})
		}
	}
	return result
}

type queueItem struct {
	point
	distance int
}

func shortestPath(matrix [][]string, start, end point) int {
	queue := []queueItem{{start, 0}}
	visited := map[point]struct{}{}
	visited[start] = struct{}{}

	for len(queue) > 0 {
		// pop first
		item := queue[0]
		queue = queue[1:]

		// reached end?
		if item.point == end {
			smallest := item.distance
			for _, q := range queue {
				if q.distance < smallest {
					smallest = q.distance
				}
			}
			return smallest
		}

		// visit neighbours
		neighbours := []point{
			{item.x - 1, item.y},
			{item.x + 1, item.y},
			{item.x, item.y - 1},
			{item.x, item.y + 1},
		}
		for _, n := range neighbours {
			inBounds := n.x >= 0 && n.x < len(matrix[0]) && n.y >= 0 && n.y < len(matrix)
			_, isVisited := visited[n]
			if inBounds && !isVisited {
				queue = append(queue, queueItem{n, item.distance + 1})
				visited[n] = struct{}{}
			}
		}
	}

	return 0
}

func parse(input string) [][]string {
	var result [][]string
	for _, row := range strings.Split(input, "\n") {
		var r []string
		for _, cell := range strings.Split(row, "") {
			r = append(r, cell)
		}
		result = append(result, r)
	}
	return result
}
