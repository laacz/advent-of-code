package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// buildSizeTree builds a map of directory -> its size
// using the input from the puzzle
func buildSizeTree(input string) map[string]int {
	var sizes = map[string]int{}

	var path []string
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "$ cd ") {
			p := line[5:]
			switch p {
			case "/":
				path = []string{}
			case "..":
				path = path[:len(path)-1]
			default:
				path = append(path, p)
			}
		} else if line[0] >= '0' && line[0] <= '9' {
			fullpath := "/" + strings.Join(path, "/")

			_, ok := sizes[fullpath]
			if !ok {
				sizes[fullpath] = 0
			}

			p := strings.Split(line, " ")
			size, _ := strconv.Atoi(p[0])
			for s := range sizes {
				if strings.HasPrefix(fullpath, s) {
					sizes[s] += size
				}
			}
		}
	}

	return sizes
}

func partOne(input string) int {
	sizes := buildSizeTree(input)
	result := 0
	for _, s := range sizes {
		if s <= 100000 {
			result += s
		}
	}

	return result
}

func partTwo(input string) int {
	sizes := buildSizeTree(input)
	to_free := 30000000 - (70000000 - sizes["/"])

	// sort the folders by size
	var keys []string
	for k := range sizes {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return sizes[keys[i]] < sizes[keys[j]]
	})

	// find the first folder that is bigger than to_free
	result := 0
	for _, k := range keys {
		if sizes[k] > to_free {
			result = sizes[k]
			break
		}
	}

	return result
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
