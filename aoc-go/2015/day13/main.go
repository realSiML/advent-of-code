package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

type Edge [3]int

func ParseLine(line string) (from string, to string, value int) {
	parts := strings.Fields(line)
	from, to = parts[0], parts[len(parts)-1]
	value, err := strconv.Atoi(parts[3])
	if err != nil {
		log.Fatal(err)
	}
	if parts[2] == "lose" {
		value = -value
	}
	return from, to[:len(to)-1], value
}

func ParseFile(path string) ([][]int, error) {
	m := map[string]int{}
	var curIndex int
	edgeList := make([]Edge, 0)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		from, to, dist := ParseLine(line)
		if _, ok := m[from]; !ok {
			m[from] = curIndex
			curIndex++
		}
		if _, ok := m[to]; !ok {
			m[to] = curIndex
			curIndex++
		}
		edgeList = append(edgeList, Edge{m[from], m[to], dist})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	n := len(m)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for _, edge := range edgeList {
		u, v, weight := edge[0], edge[1], edge[2]
		matrix[u][v] = weight
	}

	return matrix, nil
}

func Solve(dist [][]int, compare func(a, b int) int) int {
	n := len(dist)
	if n <= 1 {
		return 0
	}

	// Prepare memo
	rows, cols := 1<<n, n
	dp := make([][]int, rows)
	data := make([]int, rows*cols)
	visited := make([]bool, rows*cols)

	for i := range dp {
		dp[i] = data[i*cols : (i+1)*cols]
	}

	// Base case
	data[cols] = 0
	visited[cols] = true

	// Main logic
	for mask := 3; mask < 1<<n; mask += 2 {
		for m := mask >> 1; m > 0; m &= m - 1 {
			last := bits.TrailingZeros(uint(m)) + 1
			prevMask := mask ^ (1 << last)
			currIdx := mask*cols + last
			for pm := prevMask; pm > 0; pm &= pm - 1 {
				prev := bits.TrailingZeros(uint(pm))
				prevIdx := prevMask*cols + prev

				if !visited[prevIdx] {
					continue
				}

				newCost := data[prevIdx] + dist[prev][last] + dist[last][prev]
				if !visited[currIdx] || compare(newCost, data[currIdx]) < 0 {
					data[currIdx] = newCost
					visited[currIdx] = true
				}
			}
		}
	}

	allVisited := rows - 1
	bestCost := 0
	found := false

	for last := 1; last < n; last++ {
		idx := allVisited*cols + last
		if !visited[idx] {
			continue
		}

		cycleCost := data[idx] + dist[last][0] + dist[0][last]
		if !found || compare(cycleCost, bestCost) < 0 {
			bestCost = cycleCost
			found = true
		}
	}

	return bestCost
}

func main() {
	matrix, err := ParseFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	partOne := Solve(matrix, func(a, b int) int { return cmp.Compare(b, a) })
	fmt.Println(partOne)

	for i := range matrix {
		matrix[i] = append(matrix[i], 0)
	}
	matrix = append(matrix, make([]int, len(matrix[0])))

	partTwo := Solve(matrix, func(a, b int) int { return cmp.Compare(b, a) })
	fmt.Println(partTwo)
}
