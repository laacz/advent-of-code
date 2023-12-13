package main

import (
	"fmt"
	"os"
	"strings"
)

type Ingredient struct {
	name  string
	props map[string]int
}

func parse(line string) []Ingredient {
	ret := []Ingredient{}

	for _, line := range strings.Split(line, "\n") {
		ingredient := Ingredient{}
		ingredient.props = make(map[string]int)
		ingredient.name = strings.Split(line, ":")[0]
		for _, prop := range strings.Split(strings.Split(line, ":")[1], ",") {
			prop = strings.TrimSpace(prop)
			ingredient.props[strings.Split(prop, " ")[0]] = atoi(strings.Split(prop, " ")[1])
		}
		ret = append(ret, ingredient)
	}

	return ret
}

func atoi(s string) (ret int) {
	fmt.Sscanf(s, "%d", &ret)
	return ret
}

func combinations(total, n int) [][]int {
	ret := [][]int{}
	if n == 1 {
		ret = append(ret, []int{total})
	} else {
		for i := 1; i <= total; i++ {
			for _, combo := range combinations(total-i, n-1) {
				ret = append(ret, append([]int{i}, combo...))
			}
		}
	}
	return ret
}

func partOne(line string) (ret int) {
	ingredients := parse(line)
	for _, combo := range combinations(100, len(ingredients)) {
		props := make(map[string]int)
		for i, ingredient := range ingredients {
			for prop, val := range ingredient.props {
				if prop != "calories" {
					props[prop] += val * combo[i]
				}
			}
		}

		curr := 1
		for _, val := range props {
			if val > 0 {
				curr *= val
			}
		}

		ret = max(ret, curr)
	}
	return ret
}

func partTwo(line string) (ret int) {
	ingredients := parse(line)
	for _, combo := range combinations(100, len(ingredients)) {
		props := make(map[string]int)
		calories := 0
		for i, ingredient := range ingredients {
			for prop, val := range ingredient.props {
				if prop == "calories" {
					calories += val * combo[i]
				} else {
					props[prop] += val * combo[i]
				}
			}
		}

		if calories != 500 {
			continue
		}

		score := 1
		for _, val := range props {
			if val > 0 {
				score *= val
			}
		}

		ret = max(ret, score)

	}
	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
