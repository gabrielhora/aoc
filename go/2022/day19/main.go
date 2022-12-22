package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func main() {
	part1()
	part2()
}

func part1() {
	blueprints := parse(input)
	s := time.Now()
	ans := 0
	for _, bp := range blueprints {
		geodes := solve(bp, []state{{oreRobot: 1, minLeft: 24}}, map[state]*struct{}{})
		fmt.Printf("blueprint: %d, geodes: %v\n", bp.id, geodes)
		ans += bp.id * int(geodes)
	}
	fmt.Printf("part1: %v (%s)\n", ans, time.Now().Sub(s)) // 2m36s
}

func part2() {
	blueprints := parse(input)[:3]
	s := time.Now()
	ans := 1
	for _, bp := range blueprints {
		geodes := solve(bp, []state{{oreRobot: 1, minLeft: 32}}, map[state]*struct{}{})
		fmt.Printf("blueprint: %d, geodes: %v\n", bp.id, geodes)
		ans *= int(geodes)
	}
	fmt.Printf("part2: %v (%s)\n", ans, time.Now().Sub(s)) // 1m32s
}

func solve(bp blueprint, q []state, visited map[state]*struct{}) float64 {
	best := float64(0)
	for len(q) > 0 {
		var cur state
		cur, q = q[0], q[1:]
		if _, ok := visited[cur]; ok {
			continue
		}
		visited[cur] = nil
		// don't care about nodes with less geodes (and minutes left) than what we already found
		if cur.geode < best {
			continue
		}
		best = math.Max(best, cur.geode)
		if cur.minLeft == 0 {
			continue
		}
		// resources
		q = append(q, state{
			ore:           cur.ore + cur.oreRobot,
			clay:          cur.clay + cur.clayRobot,
			obsidian:      cur.obsidian + cur.obsidianRobot,
			geode:         cur.geode + cur.geodeRobot,
			oreRobot:      cur.oreRobot,
			clayRobot:     cur.clayRobot,
			obsidianRobot: cur.obsidianRobot,
			geodeRobot:    cur.geodeRobot,
			minLeft:       cur.minLeft - 1,
		})
		// geode robot
		if cur.ore >= bp.geodeRobot.ore && cur.obsidian >= bp.geodeRobot.obsidian {
			q = append(q, state{
				ore:           cur.ore + cur.oreRobot - bp.geodeRobot.ore,
				clay:          cur.clay + cur.clayRobot,
				obsidian:      cur.obsidian + cur.obsidianRobot - bp.geodeRobot.obsidian,
				geode:         cur.geode + cur.geodeRobot,
				oreRobot:      cur.oreRobot,
				clayRobot:     cur.clayRobot,
				obsidianRobot: cur.obsidianRobot,
				geodeRobot:    cur.geodeRobot + 1,
				minLeft:       cur.minLeft - 1,
			})
		}
		// obsidian robot
		if cur.ore >= bp.obsidianRobot.ore && cur.clay >= bp.obsidianRobot.clay {
			q = append(q, state{
				ore:           cur.ore + cur.oreRobot - bp.obsidianRobot.ore,
				clay:          cur.clay + cur.clayRobot - bp.obsidianRobot.clay,
				obsidian:      cur.obsidian + cur.obsidianRobot,
				geode:         cur.geode + cur.geodeRobot,
				oreRobot:      cur.oreRobot,
				clayRobot:     cur.clayRobot,
				obsidianRobot: cur.obsidianRobot + 1,
				geodeRobot:    cur.geodeRobot,
				minLeft:       cur.minLeft - 1,
			})
		}
		// ore robot
		if cur.ore >= bp.oreRobot.ore {
			q = append(q, state{
				ore:           cur.ore + cur.oreRobot - bp.oreRobot.ore,
				clay:          cur.clay + cur.clayRobot,
				obsidian:      cur.obsidian + cur.obsidianRobot,
				geode:         cur.geode + cur.geodeRobot,
				oreRobot:      cur.oreRobot + 1,
				clayRobot:     cur.clayRobot,
				obsidianRobot: cur.obsidianRobot,
				geodeRobot:    cur.geodeRobot,
				minLeft:       cur.minLeft - 1,
			})
		}
		// clay robot
		if cur.ore >= bp.clayRobot.ore {
			q = append(q, state{
				ore:           cur.ore + cur.oreRobot - bp.clayRobot.ore,
				clay:          cur.clay + cur.clayRobot,
				obsidian:      cur.obsidian + cur.obsidianRobot,
				geode:         cur.geode + cur.geodeRobot,
				oreRobot:      cur.oreRobot,
				clayRobot:     cur.clayRobot + 1,
				obsidianRobot: cur.obsidianRobot,
				geodeRobot:    cur.geodeRobot,
				minLeft:       cur.minLeft - 1,
			})
		}
	}
	return best
}

type state struct {
	ore, clay, obsidian, geode                     float64
	oreRobot, clayRobot, obsidianRobot, geodeRobot float64
	minLeft                                        int
}

type cost struct {
	ore, clay, obsidian float64
}

type blueprint struct {
	id                                             int
	oreRobot, clayRobot, obsidianRobot, geodeRobot cost
}

var reLine = regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)

func parse(input string) []blueprint {
	var res []blueprint
	for _, line := range strings.Split(input, "\n") {
		caps := reLine.FindStringSubmatch(line)
		id, _ := strconv.Atoi(caps[1])
		oreOre, _ := strconv.ParseFloat(caps[2], 64)
		clayOre, _ := strconv.ParseFloat(caps[3], 64)
		obsOre, _ := strconv.ParseFloat(caps[4], 64)
		obsClay, _ := strconv.ParseFloat(caps[5], 64)
		geoOre, _ := strconv.ParseFloat(caps[6], 64)
		geoObs, _ := strconv.ParseFloat(caps[7], 64)
		res = append(res, blueprint{
			id:            id,
			oreRobot:      cost{ore: oreOre},
			clayRobot:     cost{ore: clayOre},
			obsidianRobot: cost{ore: obsOre, clay: obsClay},
			geodeRobot:    cost{ore: geoOre, obsidian: geoObs},
		})
	}
	return res
}
