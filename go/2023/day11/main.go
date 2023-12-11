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
	for _, pair := range p {
		nr := newRowsBetween(image, pair[0].y, pair[1].y)
		nc := newColsBetween(image, pair[0].x, pair[1].x)
		dist := abs(pair[0].y-pair[1].y) + abs(pair[0].x-pair[1].x)
		totalDistance += dist + nr*weight + nc*weight
	}
	return totalDistance
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
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
