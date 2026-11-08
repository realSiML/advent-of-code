package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CheckNiceStringOne(s string) bool {
	vowels := make(map[rune]int)
	hasDoubleLetters := false
	window := "__"

	for _, r := range s {
		if strings.ContainsRune("aeiuo", r) {
			vowels[r] += 1
		}

		window = string(window[1]) + string(r)
		if window[0] == window[1] {
			hasDoubleLetters = true
		}

		if window == "ab" || window == "cd" || window == "pq" || window == "xy" {
			return false
		}
	}

	if !hasDoubleLetters {
		return false
	}

	totalVowels := 0
	for _, vowelCount := range vowels {
		totalVowels += vowelCount
	}

	return totalVowels >= 3
}

func PartOne() {
	niceStringCount := 0

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			line = strings.TrimRight(line, "\n")
			if CheckNiceStringOne(line) {
				niceStringCount += 1
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}

	fmt.Println(niceStringCount)
}

func CheckNiceStringTwo(s string) bool {
	pairsIndexes := make(map[string]int)
	firstCondition := false
	secondCondition := false

	for i := 1; i < len(s); i++ {
		curIndex, curPair := i-1, string(s[i-1])+string(s[i])
		_, ok := pairsIndexes[curPair]
		if !ok {
			pairsIndexes[curPair] = curIndex
		} else {
			if curIndex-1 != pairsIndexes[curPair] {
				firstCondition = true
			}
		}

		if i >= 2 && s[i-2] == s[i] {
			secondCondition = true
		}
	}

	return firstCondition && secondCondition
}

func PartTwo() {
	niceStringCount := 0

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			line = strings.TrimRight(line, "\n")
			if CheckNiceStringTwo(line) {
				niceStringCount += 1
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}

	fmt.Println(niceStringCount)
}

func main() {
	PartOne()
	PartTwo()
}
