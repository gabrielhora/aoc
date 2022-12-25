package main

import (
	_ "embed"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"hora.dev/aoc/2022/utils/math"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	data := parse(input)
	row := 2000000

	var ranges [][2]int
	for _, d := range data {
		r, err := sensorXRange(d.sensor, d.dist, row)
		if err != nil {
			continue
		}
		ranges = append(ranges, r)
	}
	merged := mergeRanges(ranges, [2]int{-1000000000, 1000000000})

	fmt.Printf("part 1: %d\n", merged[0][1]-merged[0][0])
}

func part2(input string) {
	data := parse(input)
	limit := 4000000

	x := -1
	y := -1
	for i := 0; i < limit; i++ {
		var ranges [][2]int
		for _, d := range data {
			r, err := sensorXRange(d.sensor, d.dist, i)
			if err != nil {
				continue
			}
			ranges = append(ranges, r)
		}
		merged := mergeRanges(ranges, [2]int{0, limit})[0]
		if merged[0] != 0 || merged[1] != limit {
			// found the hole
			y = i
			if merged[0] != 0 {
				x = merged[0] - 1
			} else {
				x = merged[1] + 1
			}
		}
	}

	fmt.Printf("part 2: %d\n", x*4000000+y)
}

func sensorXRange(sensor coord, dist, yRow int) ([2]int, error) {
	dy := math.Abs(yRow - sensor.y)
	if dy > dist {
		return [2]int{}, errors.New("not reachable")
	}
	return [2]int{sensor.x - (dist - dy), sensor.x + (dist - dy)}, nil
}

func mergeRanges(ranges [][2]int, maxRange [2]int) [][2]int {
	cur, tail := ranges[0], ranges[1:]
	var newRanges [][2]int
	var err error
	for _, r := range tail {
		cur, err = mergeRange(cur, r, maxRange)
		if err != nil {
			newRanges = append(newRanges, r)
		} else {
			newRanges = append(newRanges, cur)
		}
	}

	if len(newRanges) > 0 && len(newRanges) < len(ranges) {
		return mergeRanges(newRanges, maxRange)
	}

	return ranges
}

func mergeRange(a, b, maxRange [2]int) ([2]int, error) {
	a1 := math.Max(a[0], maxRange[0])
	a2 := math.Min(a[1], maxRange[1])
	b1 := math.Max(b[0], maxRange[0])
	b2 := math.Min(b[1], maxRange[1])

	// overlap
	if a1 <= b2 && a2 >= b1 {
		return [2]int{math.Min(a1, b1), math.Max(a2, b2)}, nil
	}
	// ends touching
	if a2 < b1 && math.Abs(a2-b1) == 1 {
		return [2]int{a1, b2}, nil
	}
	if b2 < a1 && math.Abs(b2-a1) == 1 {
		return [2]int{b1, a2}, nil
	}
	// no overlap
	return a, errors.New("no overlap")
}

var sensorRe = regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

type coord struct {
	x, y int
}

type sensorAndBeacon struct {
	sensor, beacon coord
	dist           int
}

func parse(input string) []sensorAndBeacon {
	var res []sensorAndBeacon
	for _, l := range strings.Split(input, "\n") {
		caps := sensorRe.FindStringSubmatch(l)
		sx, _ := strconv.Atoi(caps[1])
		sy, _ := strconv.Atoi(caps[2])
		bx, _ := strconv.Atoi(caps[3])
		by, _ := strconv.Atoi(caps[4])
		sensor := coord{sx, sy}
		beacon := coord{bx, by}
		res = append(res, sensorAndBeacon{
			sensor: sensor,
			beacon: beacon,
			dist:   dist(sensor, beacon),
		})
	}
	return res
}

func dist(a, b coord) int {
	return math.Abs(a.x-b.x) + math.Abs(a.y-b.y)
}
