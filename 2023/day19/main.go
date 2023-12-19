package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Workflow struct {
	rules []Rule
}

func (w Workflow) Run(p Part) string {
	for _, rule := range w.rules {
		if rule.Test(p) {
			return rule.target
		}
	}
	return ""
}

type Rule struct {
	category, test string
	value          int
	target         string
}

func (r Rule) Test(part Part) bool {
	if r.category == "" {
		return true
	}

	var val int
	switch r.category {
	case "x":
		val = part.x
	case "m":
		val = part.m
	case "a":
		val = part.a
	case "s":
		val = part.s
	}

	switch r.test {
	case "<":
		return val < r.value
	case ">":
		return val > r.value
	}

	return false
}

type Workflows map[string]Workflow

type Part struct {
	x, m, a, s int
}

// parse parses the input into a Grid according to the rules of the first part of the puzzle
func parse(input string) (Workflows, []Part) {
	blocks := strings.Split(input, "\n\n")
	wf := make(Workflows)

	re := regexp.MustCompile(`([a-z]+)([<>])([0-9]+):([a-zA-Z]+)`)
	for _, line := range strings.Split(blocks[0], "\n") {
		name := strings.Split(line, "{")[0]
		rules := strings.Split(line, "{")[1]
		rules = strings.TrimSuffix(rules, "}")
		workflow := Workflow{}
		for _, rule := range strings.Split(rules, ",") {
			r := Rule{}
			if re.MatchString(rule) {
				r.category = re.FindStringSubmatch(rule)[1]
				r.test = re.FindStringSubmatch(rule)[2]
				fmt.Sscanf(re.FindStringSubmatch(rule)[3], "%d", &r.value)
				r.target = re.FindStringSubmatch(rule)[4]
			} else {
				r.target = rule
			}
			workflow.rules = append(workflow.rules, r)
		}
		wf[name] = workflow
	}

	var parts []Part
	for _, line := range strings.Split(blocks[1], "\n") {
		line = strings.TrimFunc(line, func(r rune) bool {
			return r == '{' || r == '}'
		})
		var part Part

		for _, attr := range strings.Split(line, ",") {
			var a rune
			var val int

			fmt.Sscanf(attr, "%c=%d", &a, &val)
			switch a {
			case 'x':
				part.x = val
			case 'm':
				part.m = val
			case 'a':
				part.a = val
			case 's':
				part.s = val
			}

		}
		parts = append(parts, part)
	}

	return wf, parts
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) (ret int) {
	wf, parts := parse(input)

	for _, part := range parts {
		target := "in"
		for {
			workflow := wf[target]
			target = workflow.Run(part)
			if target == "A" || target == "R" || target == "" {
				if target == "A" {
					ret += part.a + part.m + part.s + part.x
				}
				break
			}
		}
	}
	return ret
}

type Range struct {
	xmin, xmax, mmin, mmax, amin, amax, smin, smax int
}

func (wfs Workflows) getRanges(wf string, ra Range) (ret []Range) {
	r := ra

	if wf == "R" {
		return ret
	}

	if wf == "A" {
		return []Range{r}
	}

	for _, rule := range wfs[wf].rules {
		switch rule.test {
		case "<":
			rr := r
			setmaxval(&rr, rule.category, rule.value-1)
			ret = append(ret, wfs.getRanges(rule.target, rr)...)
			setminval(&r, rule.category, rule.value)
		case ">":
			rr := r
			setminval(&rr, rule.category, rule.value+1)
			ret = append(ret, wfs.getRanges(rule.target, rr)...)
			setmaxval(&r, rule.category, rule.value)
		default:
			ret = append(ret, wfs.getRanges(rule.target, r)...)
		}
	}

	return ret
}

func setminval(r *Range, category string, value int) {
	switch category {
	case "x":
		r.xmin = value
	case "m":
		r.mmin = value
	case "a":
		r.amin = value
	case "s":
		r.smin = value
	}
}

func setmaxval(r *Range, category string, value int) {
	switch category {
	case "x":
		r.xmax = value
	case "m":
		r.mmax = value
	case "a":
		r.amax = value
	case "s":
		r.smax = value
	}
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) (ret int) {

	wfs, _ := parse(input)

	ranges := wfs.getRanges("in", Range{1, 4000, 1, 4000, 1, 4000, 1, 4000})

	for _, r := range ranges {
		ret += (r.xmax - r.xmin + 1) * (r.mmax - r.mmin + 1) * (r.amax - r.amin + 1) * (r.smax - r.smin + 1)
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
