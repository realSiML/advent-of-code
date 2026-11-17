package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var m = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func ParseFile(path string) []map[string]int {
	aunts := make([]map[string]int, 0)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		aunts = append(aunts, ParseLine(sc.Text()))
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return aunts
}

func ParseLine(line string) map[string]int {
	m := map[string]int{}

	parts := strings.FieldsFunc(line, func(r rune) bool {
		return r == ' ' || r == ':' || r == ','
	})

	num, _ := strconv.Atoi(parts[3])
	m[parts[2]] = num
	num, _ = strconv.Atoi(parts[5])
	m[parts[4]] = num
	num, _ = strconv.Atoi(parts[7])
	m[parts[6]] = num

	return m
}

func PartOne(aunts []map[string]int) int {
	for i, aunt := range aunts {
		found := true
		for k, got := range aunt {
			if want := m[k]; got != want {
				found = false
				break
			}
		}
		if found {
			return i + 1
		}
	}

	return 0
}

func PartTwo(aunts []map[string]int) int {
	for i, aunt := range aunts {
		found := true
	OuterLoop:
		for k, got := range aunt {
			want := m[k]
			switch k {
			case "cats", "trees":
				if got <= want {
					found = false
					break OuterLoop
				}
			case "pomeranians", "goldfish":
				if got >= want {
					found = false
					break OuterLoop
				}
			default:
				if got != want {
					found = false
					break OuterLoop
				}
			}
		}
		if found {
			return i + 1
		}
	}

	return 0
}

func main() {
	aunts := ParseFile("./input.txt")

	partOne := PartOne(aunts)
	fmt.Println(partOne)

	partTwo := PartTwo(aunts)
	fmt.Println(partTwo)
}
