package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	seeds, seedPath := parse(input)
	result := int64(math.MaxInt64)
	for _, s := range seeds {
		// walk the seed "path", last number is the location
		for _, p := range seedPath {
			s = p.get(s)
		}
		if s < result {
			result = s
		}
	}
	fmt.Printf("part 1: %v\n", result)
}

func part2(input string) {
	seedPairs, seedPath := parse(input)

	var seedRanges []intRange
	for i := 0; i < len(seedPairs); i += 2 {
		start := seedPairs[i]
		length := seedPairs[i+1]
		seedRanges = append(seedRanges, intRange{src: start, dst: start, length: length})
	}

	locations := make(chan int64, len(seedRanges))
	wg := sync.WaitGroup{}

	// process all ranges in parallel
	for _, r := range seedRanges {
		wg.Add(1)

		go func(r intRange, wg *sync.WaitGroup) {
			defer wg.Done()

			rangeSeed := r.src
			count := 0.0
			smallest := int64(math.MaxInt64)

			for {
				// show progress
				if math.Mod(float64(rangeSeed), 1000000) == 0 {
					count += 1
					pct := count / (float64(r.length) / 1000000) * 100
					fmt.Printf("start = %d, current = %.2f\n", r.src, pct)
				}

				// walk the seed "path" and get the location
				pathSeed := rangeSeed
				for _, p := range seedPath {
					pathSeed = p.get(pathSeed)
				}
				if pathSeed < smallest {
					smallest = pathSeed
				}

				// break out when we finished with the range
				rangeSeed = r.get(rangeSeed + 1)
				if rangeSeed == -1 {
					break
				}
			}

			locations <- smallest
		}(r, &wg)
	}

	go func() {
		wg.Wait()
		close(locations)
	}()

	// get the smallest location of all paths
	result := int64(math.MaxInt64)
	for l := range locations {
		if l < result {
			result = l
		}
	}

	fmt.Printf("part 2: %v\n", result)
}

func parse(input string) ([]int64, []seedMap) {
	maps := strings.Split(input, "\n\n")

	seedsS := strings.Split(maps[0][7:], " ")
	var seeds []int64
	for _, s := range seedsS {
		n, _ := strconv.ParseInt(s, 10, 64)
		seeds = append(seeds, n)
	}

	seedToSoil := parseMap(maps[1])
	soilToFertilizer := parseMap(maps[2])
	fertilizerToWater := parseMap(maps[3])
	waterToLight := parseMap(maps[4])
	lightToTemperature := parseMap(maps[5])
	temperatureToHumidity := parseMap(maps[6])
	humidityToLocation := parseMap(maps[7])

	return seeds, []seedMap{
		seedToSoil,
		soilToFertilizer,
		fertilizerToWater,
		waterToLight,
		lightToTemperature,
		temperatureToHumidity,
		humidityToLocation,
	}
}

func parseMap(m string) seedMap {
	var maps []intRange
	for _, line := range strings.Split(m, "\n")[1:] {
		parts := strings.Split(line, " ")
		dst, _ := strconv.ParseInt(parts[0], 10, 64)
		src, _ := strconv.ParseInt(parts[1], 10, 64)
		length, _ := strconv.ParseInt(parts[2], 10, 64)
		maps = append(maps, intRange{src, dst, length})
	}
	return seedMap{maps}
}

type seedMap struct {
	maps []intRange
}

func (sm seedMap) get(at int64) int64 {
	for _, m := range sm.maps {
		if d := m.get(at); d != -1 {
			return d
		}
	}
	return at
}

type intRange struct {
	src, dst, length int64
}

func (r intRange) get(at int64) int64 {
	if at >= r.src && at < r.src+r.length {
		return r.dst + at - r.src
	} else {
		return -1
	}
}
