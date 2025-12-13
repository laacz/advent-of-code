package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Content struct {
	Color string
	Count int
}

func parse(input string) map[string][]Content {
	rules := map[string][]Content{}
	re := regexp.MustCompile(`(\d+) (\w+ \w+) bags?`)

	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, " bags contain ")
		outer := parts[0]
		rules[outer] = []Content{}

		if strings.Contains(parts[1], "no other bags") {
			continue
		}

		for _, match := range re.FindAllStringSubmatch(parts[1], -1) {
			count, _ := strconv.Atoi(match[1])
			rules[outer] = append(rules[outer], Content{Color: match[2], Count: count})
		}
	}

	return rules
}

func part1(rules map[string][]Content) int {
	reverse := map[string][]string{}
	for parent, contents := range rules {
		for _, c := range contents {
			reverse[c.Color] = append(reverse[c.Color], parent)
		}
	}

	visited := map[string]bool{}
	queue := []string{"shiny gold"}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, parent := range reverse[cur] {
			if !visited[parent] {
				visited[parent] = true
				queue = append(queue, parent)
			}
		}
	}

	return len(visited)
}

func part2(rules map[string][]Content) int {
	var count func(color string) int
	count = func(color string) int {
		total := 0
		for _, c := range rules[color] {
			total += c.Count * (1 + count(c.Color))
		}
		return total
	}
	return count("shiny gold")
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1:", part1(parse(string(data))))
	fmt.Println("Part 2:", part2(parse(string(data))))
}
