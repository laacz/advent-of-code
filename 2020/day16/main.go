package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Range struct {
	min, max int
}

type Rule struct {
	name   string
	ranges []Range
}

func (r Rule) isValid(n int) bool {
	for _, rng := range r.ranges {
		if n >= rng.min && n <= rng.max {
			return true
		}
	}
	return false
}

type Input struct {
	rules         []Rule
	myTicket      []int
	nearbyTickets [][]int
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := part1(string(data))
	fmt.Println("Part 1:", result)

	result2 := part2(string(data))
	fmt.Println("Part 2:", result2)
}

func parse(input string) Input {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")

	var rules []Rule
	ruleRe := regexp.MustCompile(`^(.+): (\d+)-(\d+) or (\d+)-(\d+)$`)
	for _, line := range strings.Split(parts[0], "\n") {
		matches := ruleRe.FindStringSubmatch(line)
		if matches != nil {
			min1, _ := strconv.Atoi(matches[2])
			max1, _ := strconv.Atoi(matches[3])
			min2, _ := strconv.Atoi(matches[4])
			max2, _ := strconv.Atoi(matches[5])
			rules = append(rules, Rule{
				name:   matches[1],
				ranges: []Range{{min1, max1}, {min2, max2}},
			})
		}
	}

	myTicketLine := strings.Split(parts[1], "\n")[1]
	myTicket := parseTicket(myTicketLine)

	var nearbyTickets [][]int
	for _, line := range strings.Split(parts[2], "\n")[1:] {
		nearbyTickets = append(nearbyTickets, parseTicket(line))
	}

	return Input{rules, myTicket, nearbyTickets}
}

func parseTicket(line string) []int {
	var ticket []int
	for _, numStr := range strings.Split(line, ",") {
		num, _ := strconv.Atoi(numStr)
		ticket = append(ticket, num)
	}
	return ticket
}

func part1(input string) int {
	data := parse(input)

	errorRate := 0
	for _, ticket := range data.nearbyTickets {
		for _, num := range ticket {
			validForAny := false
			for _, rule := range data.rules {
				if rule.isValid(num) {
					validForAny = true
					break
				}
			}
			if !validForAny {
				errorRate += num
			}
		}
	}

	return errorRate
}

func part2(input string) int {
	data := parse(input)

	var validTickets [][]int
	for _, ticket := range data.nearbyTickets {
		valid := true
		for _, num := range ticket {
			validForAny := false
			for _, rule := range data.rules {
				if rule.isValid(num) {
					validForAny = true
					break
				}
			}
			if !validForAny {
				valid = false
				break
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}

	numFields := len(data.myTicket)
	possible := make([]map[int]bool, numFields)
	for i := 0; i < numFields; i++ {
		possible[i] = make(map[int]bool)
		for ruleIdx := range data.rules {
			possible[i][ruleIdx] = true
		}
	}

	for _, ticket := range validTickets {
		for pos, val := range ticket {
			for ruleIdx := range possible[pos] {
				if !data.rules[ruleIdx].isValid(val) {
					delete(possible[pos], ruleIdx)
				}
			}
		}
	}

	fieldMapping := make(map[int]int)
	for len(fieldMapping) < numFields {
		for pos, rules := range possible {
			if len(rules) == 1 {
				for ruleIdx := range rules {
					fieldMapping[pos] = ruleIdx

					for i := range possible {
						delete(possible[i], ruleIdx)
					}
				}
			}
		}
	}

	result := 1
	for pos, ruleIdx := range fieldMapping {
		if strings.HasPrefix(data.rules[ruleIdx].name, "departure") {
			result *= data.myTicket[pos]
		}
	}

	return result
}
