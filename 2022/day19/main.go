package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Ore int = iota
	Clay
	Obsidian
	Geode
)

type Blueprint [][]int

func Parse(input string) []Blueprint {
	var ret []Blueprint

	for _, line := range strings.Split(input, "\n") {
		var bp, o, c, o1, o2, g1, g2 int
		_, _ = fmt.Sscanf(
			line,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&bp, &o, &c, &o1, &o2, &g1, &g2,
		)

		blueprint := Blueprint{
			{o, 0, 0, 0},
			{c, 0, 0, 0},
			{o1, o2, 0, 0},
			{g1, 0, g2, 0},
		}

		ret = append(ret, blueprint)
	}

	return ret
}

type State struct {
	inventory []int // [ore, clay, obsidian, geode] items
	robots    []int // [ore, clay, obsidian, geode] robots
	time      int
}

func maxGeodesBFS(blueprint Blueprint, maxTime int) int {
	maxGeodes := 0

	q := []*State{
		{
			inventory: []int{0, 0, 0, 0},
			robots:    []int{1, 0, 0, 0},
			time:      0,
		},
	}

	maxRobots := []int{1000, 1000, 1000, 1000}
	for i := 0; i < 3; i++ {
		for j := 0; j < len(blueprint); j++ {
			if blueprint[j][i] > 0 {
				maxRobots[i] = blueprint[j][i]
			}
		}
	}

	// Fastest time to the first geode (for pruning)
	bestGeodes := make(map[int]int)
	for len(q) != 0 {
		s := q[0]
		q = q[1:]
		for i := 0; i < len(blueprint); i++ {
			if s.robots[i] == maxRobots[i]+1 {
				continue
			}
			costs := blueprint[i]
			waitTime := 0
			for j := 0; j < 3; j++ {
				wt := 0
				if s.inventory[j] >= costs[j] {
					wt = 0
				} else if s.robots[j] == 0 {
					wt = maxTime + 1
				} else {
					wt = (costs[j] - s.inventory[j] + s.robots[j] - 1) / s.robots[j]
				}
				if wt > waitTime {
					waitTime = wt
				}
			}

			newTime := s.time + waitTime + 1
			if newTime >= maxTime {
				continue
			}

			newInventory := []int{
				s.inventory[Ore] + s.robots[Ore]*(waitTime+1) - costs[Ore],
				s.inventory[Clay] + s.robots[Clay]*(waitTime+1) - costs[Clay],
				s.inventory[Obsidian] + s.robots[Obsidian]*(waitTime+1) - costs[Obsidian],
				s.inventory[Geode] + s.robots[Geode]*(waitTime+1) - costs[Geode],
			}

			// Prune branches where the amount of geodes ir less than for the best performer
			if best, ok := bestGeodes[newTime]; ok && newInventory[Geode] < best-1 {
				continue
			}
			bestGeodes[newTime] = newInventory[Geode]

			var newRobots = []int{
				s.robots[Ore],
				s.robots[Clay],
				s.robots[Obsidian],
				s.robots[Geode],
			}

			newRobots[i] += 1

			q = append(q, &State{
				inventory: newInventory,
				robots:    newRobots,
				time:      newTime,
			})
		}

		geodes := s.inventory[Geode] + s.robots[Geode]*(maxTime-s.time)
		if maxGeodes < geodes {
			maxGeodes = geodes
		}
	}

	return maxGeodes
}

func partOne(input string) int {
	ret := 0
	p := Parse(input)
	for i, bp := range p {
		mg := maxGeodesBFS(bp, 24)
		ret += mg * (i + 1)
	}

	return ret
}

func partTwo(input string) int {
	ret := 1
	p := Parse(input)
	for i, bp := range p {
		if i > 2 {
			break
		}
		mg := maxGeodesBFS(bp, 32)
		ret *= mg
	}

	return ret
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
