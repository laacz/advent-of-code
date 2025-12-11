package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(input string) []map[string]string {
	ret := []map[string]string{}
	groups := strings.Split(strings.TrimSpace(input), "\n\n")
	for _, group := range groups {
		passport := map[string]string{}
		fields := strings.Fields(group)
		for _, field := range fields {
			parts := strings.SplitN(field, ":", 2)
			passport[parts[0]] = parts[1]
		}
		ret = append(ret, passport)
	}

	return ret
}

func part1(input []map[string]string) int {
	ret := 0

	for _, passport := range input {
		if passport["byr"] == "" || passport["iyr"] == "" || passport["eyr"] == "" || passport["hgt"] == "" || passport["hcl"] == "" || passport["ecl"] == "" || passport["pid"] == "" {
			continue
		}
		ret++
	}

	return ret
}

func part2(input []map[string]string) int {
	ret := 0

	for _, passport := range input {
		// byr (Birth Year) - four digits; at least 1920 and at most 2002.
		if len(passport["byr"]) != 4 || passport["byr"] < "1920" || passport["byr"] > "2002" {
			continue
		}

		// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		if len(passport["iyr"]) != 4 || passport["iyr"] < "2010" || passport["iyr"] > "2020" {
			continue
		}

		// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		if len(passport["eyr"]) != 4 || passport["eyr"] < "2020" || passport["eyr"] > "2030" {
			continue
		}

		// hgt (Height) - a number followed by either cm or in:
		// - If cm, the number must be at least 150 and at most 193.
		// - If in, the number must be at least 59 and at most 76.
		hgt := passport["hgt"]
		if strings.HasSuffix(hgt, "cm") && len(hgt) >= 4 {
			if hgt < "150cm" || hgt > "193cm" {
				continue
			}
		} else if strings.HasSuffix(hgt, "in") && len(hgt) >= 3 {
			if hgt < "59in" || hgt > "76in" {
				continue
			}
		} else {
			continue
		}

		// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		hcl := passport["hcl"]
		if len(hcl) != 7 || hcl[0] != '#' {
			continue
		}
		for _, c := range hcl[1:] {
			if (c < '0' || c > '9') && (c < 'a' || c > 'f') {
				continue
			}
		}

		// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		switch passport["ecl"] {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		default:
			continue
		}

		// pid (Passport ID) - a nine-digit number, including leading zeroes.
		if len(passport["pid"]) != 9 {
			continue
		}
		if _, err := strconv.Atoi(passport["pid"]); err != nil {
			continue
		}

		ret++
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1:", part1(parse(string(data))))
	fmt.Println("Part 2:", part2(parse(string(data))))
}
