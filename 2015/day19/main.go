package main

import (
	"fmt"
	"os"
	"strings"
)

func parse(lines string) (rules map[string][]string, molecule string) {
	rules = make(map[string][]string)
	for _, line := range strings.Split(lines, "\n") {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.Contains(line, "=>") {
			parts := strings.Split(line, " => ")
			rules[parts[0]] = append(rules[parts[0]], parts[1])
		} else {
			molecule = line
		}
	}

	return rules, molecule
}

func combinations(rules map[string][]string, molecule string) (ret map[string]string) {
	ret = make(map[string]string)

	for rule, replacements := range rules {
		for _, replacement := range replacements {
			for i := 0; i < len(molecule); i++ {
				if len(molecule[i:]) >= len(rule) && molecule[i:i+len(rule)] == rule {
					str := molecule[:i] + replacement + molecule[i+len(rule):]
					ret[str] = str
				}
			}
		}
	}

	return ret
}

func partOne(lines string) (ret int) {
	rules, mol := parse(lines)
	combos := combinations(rules, mol)
	return len(combos)
}

// fuck me, i have no clue why this works
// and it does it only one every nth run
// prolly' smthn to do with the order
func partTwo(lines string) (ret int) {
	rules, mol := parse(lines)

	newrules := make(map[string]string)
	for rule, replacements := range rules {
		for _, replacement := range replacements {
			newrules[replacement] = rule
		}
	}

	froms := []string{}
	for from := range newrules {
		froms = append(froms, from)
	}

	for mol != "e" {
		for _, r := range froms {
			if strings.Contains(mol, r) {
				ret += 1
				mol = strings.Replace(mol, r, newrules[r], 1)
			}
		}

		if mol == strings.Repeat("e", len(mol)) {
			break
		}
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
