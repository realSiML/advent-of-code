package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseSizes(line string) (int, int, int) {
	var results [3]int
	sizes := strings.Split(line, "x")

	for i, v := range sizes {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		results[i] = num
	}

	return results[0], results[1], results[2]
}

func calculateBoxArea(l, w, h int) int {
	return 2*l*w + 2*w*h + 2*l*h
}

func calculateMinSideArea(l, w, h int) int {
	sum := l + w + h
	return min(l, w, h) * (sum - min(l, w, h) - max(l, w, h))
}

func CalculateRequiredWrapping(l, w, h int) int {
	return calculateBoxArea(l, w, h) + calculateMinSideArea(l, w, h)
}

func PartOne() {
	var total int

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var l, w, h int

	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			line = strings.TrimRight(line, "\n")
			l, w, h = parseSizes(line)
			total += CalculateRequiredWrapping(l, w, h)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}

	fmt.Println(total)
}

func CalculateRequiredRibbon(l, w, h int) int {
	minSize := min(l, w, h)
	secondMinSize := (l + w + h) - min(l, w, h) - max(l, w, h)
	return 2*(minSize+secondMinSize) + l*w*h
}

func PartTwo() {
	var total int

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var l, w, h int

	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			line = strings.TrimRight(line, "\n")
			l, w, h = parseSizes(line)
			total += CalculateRequiredRibbon(l, w, h)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}

	fmt.Println(total)
}

func main() {
	PartOne()
	PartTwo()
}
