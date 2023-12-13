package main

import (
	"fmt"
	"os"
	"strings"
)

type Reindeer struct {
	name   string
	speed  int
	fly    int
	rest   int
	points int
}

func (r *Reindeer) Distance(time int) (ret int) {
	a := time / (r.fly + r.rest)

	ret = a * r.fly * r.speed

	// fmt.Print(r.name, " flies ", a, " times, reaches ", ret, " km ")

	timeleft := time % (r.fly + r.rest)

	more := 0
	if timeleft < r.fly {
		more = r.speed * (time % (r.fly + r.rest))
	} else {
		more = r.speed * r.fly
	}
	// fmt.Print("and then flies ", more, " km until time runs out")
	ret += more

	// fmt.Println()

	return ret
}

func parse(line string) (ret []Reindeer) {
	for _, line := range strings.Split(line, "\n") {
		if line == "" {
			continue
		}

		var name string
		var speed, fly, rest int

		fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &fly, &rest)
		ret = append(ret, Reindeer{name, speed, fly, rest, 0})
	}

	return ret
}

func partOne(line string, duration int) (ret int) {
	deers := parse(line)

	for _, deer := range deers {
		dist := deer.Distance(duration)
		if dist > ret {
			ret = dist
		}

	}

	return ret
}

func partTwo(line string, duration int) (ret int) {
	deers := parse(line)

	points := make(map[string]int)

	maxdistance := 0
	for i := 1; i < duration; i++ {
		for _, deer := range deers {
			distance := deer.Distance(i)
			if distance > maxdistance {
				maxdistance = distance
			}
		}

		for _, deer := range deers {
			if deer.Distance(i) == maxdistance {
				points[deer.name]++
			}
		}
	}

	for _, point := range points {
		if point > ret {
			ret = point
		}
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data), 2503))
	fmt.Printf("Part two: %d\n", partTwo(string(data), 2503))
}
