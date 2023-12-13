package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

func comparePackets(left, right interface{}) float64 {
	switch {
	case reflect.TypeOf(left) == reflect.TypeOf(0.0) &&
		reflect.TypeOf(right) == reflect.TypeOf(0.0):
		return left.(float64) - right.(float64)
	case reflect.TypeOf(left) == reflect.TypeOf(0.0):
		left = []interface{}{left}
	case reflect.TypeOf(right) == reflect.TypeOf(0.0):
		right = []interface{}{right}
	}

	leftSlice := left.([]interface{})
	rightSlice := right.([]interface{})
	for i := 0; i < len(leftSlice) && i < len(rightSlice); i++ {
		if res := comparePackets(leftSlice[i], rightSlice[i]); res != 0 {
			return res
		}
	}

	return float64(len(leftSlice) - len(rightSlice))
}

func partOne(input string) int {
	ret := 0

	var first, second interface{}
	for i, pair := range strings.Split(input, "\n\n") {
		parts := strings.Split(pair, "\n")

		_ = json.Unmarshal([]byte(parts[0]), &first)
		_ = json.Unmarshal([]byte(parts[1]), &second)

		if comparePackets(first, second) < 0 {
			ret += i + 1
		}

	}

	return ret
}

func partTwo(input string) int {
	ret := 1

	var first, second interface{}
	var packets []interface{}
	for _, pair := range strings.Split(input, "\n\n") {
		parts := strings.Split(pair, "\n")

		_ = json.Unmarshal([]byte(parts[0]), &first)
		_ = json.Unmarshal([]byte(parts[1]), &second)

		packets = append(packets, first, second)
	}

	packets = append(packets, []interface{}{2.0}, []interface{}{6.0})

	sort.Slice(packets, func(i, j int) bool {
		return comparePackets(packets[i], packets[j]) < 0
	})

	for i, packet := range packets {
		if fmt.Sprintf("%v", packet) == "[2]" || fmt.Sprintf("%v", packet) == "[6]" {
			ret *= i + 1
		}
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
