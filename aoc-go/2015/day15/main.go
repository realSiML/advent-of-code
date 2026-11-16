package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Properties [5]int

func ParseFile(path string) []Properties {
	ingredients := make([]Properties, 0)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		ingredients = append(ingredients, ParseLine(sc.Text()))
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return ingredients
}

func ParseLine(line string) Properties {
	parts := strings.FieldsFunc(line, func(r rune) bool {
		return r == ' ' || r == ',' || r == ':'
	})

	capacity, _ := strconv.Atoi(parts[2])
	durability, _ := strconv.Atoi(parts[4])
	flavor, _ := strconv.Atoi(parts[6])
	texture, _ := strconv.Atoi(parts[8])
	calories, _ := strconv.Atoi(parts[10])

	return Properties{capacity, durability, flavor, texture, calories}
}

func CalculateScore(ingredients []Properties, amounts []int, calories int) int {
	if calories > 0 {
		c := 0
		for i := range ingredients {
			c += amounts[i] * ingredients[i][4]
		}

		if c != calories {
			return 0
		}
	}

	score := 1

	for p := range 4 {
		sum := 0
		for i := range ingredients {
			sum += amounts[i] * ingredients[i][p]
		}

		if sum < 0 {
			return 0
		}

		score *= sum
	}

	return score
}

func Solve(ingredients []Properties, teaspoons, calories int) int {
	amounts := make([]int, len(ingredients))
	bestScore := 0

	var enumerate func(idx, remaining int)
	enumerate = func(idx, remaining int) {
		if idx == len(ingredients)-1 {
			amounts[idx] = remaining
			score := CalculateScore(ingredients, amounts, calories)
			if score > bestScore {
				bestScore = score
			}
			return
		}

		for qty := 0; qty <= remaining; qty++ {
			amounts[idx] = qty
			enumerate(idx+1, remaining-qty)
		}
	}

	enumerate(0, teaspoons)

	return bestScore
}

func main() {
	const TEASPOONS = 100
	const CALORIES = 500

	ingredients := ParseFile("./input.txt")

	partOne := Solve(ingredients, TEASPOONS, 0)
	fmt.Println(partOne)

	partTwo := Solve(ingredients, TEASPOONS, CALORIES)
	fmt.Println(partTwo)
}
