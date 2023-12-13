package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func sum(line string) (ret int) {
	quotes := false
	prev := ""
	for _, c := range line {
		if c == '"' {
			quotes = !quotes
		}

		if !quotes && ((c >= '0' && c <= '9') || c == '-') {
			prev += string(c)
		}

		if !quotes && (c < '0' || c > '9') && c != '-' {
			if prev != "" {
				num := 0
				fmt.Sscanf(prev, "%d", &num)
				ret += num
				prev = ""
			}
		}
	}

	return ret
}

// fuuuuuuuu....
func cleanup(j interface{}) interface{} {
	switch m := j.(type) {
	case map[string]interface{}:
		for k, v := range m {
			if v == "red" {
				return map[string]interface{}{}
			}
			m[k] = cleanup(v)
		}
		return m
	case []interface{}:
		for i, v := range m {
			m[i] = cleanup(v)
		}
		return m
	default:
		return j
	}
}

func partOne(line string) int {
	return sum(line)
}

func partTwo(line string) (ret int) {
	var j interface{}

	json.Unmarshal([]byte(line), &j)
	j = cleanup(j)
	b, _ := json.Marshal(j)

	return sum(string(b))
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
