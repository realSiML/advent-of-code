package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Reindeer struct {
	speed        int
	flyDuration  int
	restDuration int
}

func ParseFile(path string) []Reindeer {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	reindeers := make([]Reindeer, 0)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		reindeers = append(reindeers, ParseLine(sc.Text()))
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return reindeers
}

func ParseLine(line string) Reindeer {
	parts := strings.Fields(line)

	speed, _ := strconv.Atoi(parts[3])
	flyDuration, _ := strconv.Atoi(parts[6])
	restDuration, _ := strconv.Atoi(parts[len(parts)-2])

	return Reindeer{speed, flyDuration, restDuration}
}

func CalculateDistanceTravelled(r Reindeer, seconds int) int {
	speed, flyDuration, restDuration := r.speed, r.flyDuration, r.restDuration

	cycleDuration := flyDuration + restDuration
	fullCycles := seconds / cycleDuration
	overtime := seconds % cycleDuration

	return speed * (fullCycles*flyDuration + min(overtime, flyDuration))
}

func CalculatePoints(rs []Reindeer, seconds int) []int {
	points := make([]int, len(rs))

	states := make(
		[]struct {
			distance int
			phase    int
		},
		len(rs),
	)

	for s := 1; s <= seconds; s++ {
		maxDistance := 0
		leaders := make([]int, 0, len(rs))
		for i, r := range rs {
			s := &states[i]
			if s.phase < r.flyDuration {
				s.distance += r.speed
			}
			s.phase++
			if s.phase == r.flyDuration+r.restDuration {
				s.phase = 0
			}
			if s.distance > maxDistance {
				maxDistance = s.distance
				leaders = leaders[:0]
				leaders = append(leaders, i)
			} else if s.distance == maxDistance {
				leaders = append(leaders, i)
			}
		}
		for _, i := range leaders {
			points[i]++
		}
	}

	return points
}

func main() {
	const INPUT int = 2503

	reindeers := ParseFile("./input.txt")

	var maxDistance int
	for _, r := range reindeers {
		d := CalculateDistanceTravelled(r, INPUT)
		maxDistance = max(maxDistance, d)
	}

	partOne := maxDistance
	fmt.Println(partOne)

	points := CalculatePoints(reindeers, INPUT)
	partTwo := slices.Max(points)
	fmt.Println(partTwo)
}
