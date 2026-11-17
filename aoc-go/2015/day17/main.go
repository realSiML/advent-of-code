package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func ParseFile(path string) []int {
	nums := make([]int, 0)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		num, _ := strconv.Atoi(sc.Text())
		nums = append(nums, num)
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return nums
}

func PartOne(liters int, containers []int) int {
	dp := make([]int, liters+1)
	dp[0] = 1

	for _, c := range containers {
		for i := liters; i > c-1; i-- {
			dp[i] += dp[i-c]
		}
	}

	return dp[liters]
}

func PartTwo(liters int, containers []int) int {
	const INF int = math.MaxInt - 1

	dp := make([]int, liters+1)
	ways := make([]int, liters+1)

	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0
	ways[0] = 1

	for _, c := range containers {
		for i := liters; i > c-1; i-- {
			if dp[i-c]+1 < dp[i] {
				dp[i] = dp[i-c] + 1
				ways[i] = ways[i-c]
			} else if dp[i-c]+1 == dp[i] {
				ways[i] += ways[i-c]
			}
		}
	}

	if dp[liters] == INF {
		return 0
	}

	return ways[liters]
}

func main() {
	const LITERS int = 150
	containers := ParseFile("./input.txt")

	partOne := PartOne(LITERS, containers)
	fmt.Println(partOne)

	partTwo := PartTwo(LITERS, containers)
	fmt.Println(partTwo)
}
