package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"math"
	"math/bits"
	"os"
	"slices"
	"strconv"
	"strings"
)

const INF = math.MaxInt

type Edge [3]int

func ParseLine(line string) (cityFrom string, cityTo string, dist int) {
	parts := strings.Fields(line)
	cityFrom, cityTo = parts[0], parts[2]
	dist, err := strconv.Atoi(parts[4])
	if err != nil {
		log.Fatal(err)
	}
	return cityFrom, cityTo, dist
}

func ParseFile(path string) ([][]int, error) {
	cities := map[string]int{}
	var curIndex int
	edgeList := make([]Edge, 0)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cityFrom, cityTo, dist := ParseLine(line)
		if _, ok := cities[cityFrom]; !ok {
			cities[cityFrom] = curIndex
			curIndex++
		}
		if _, ok := cities[cityTo]; !ok {
			cities[cityTo] = curIndex
			curIndex++
		}
		edgeList = append(edgeList, Edge{cities[cityFrom], cities[cityTo], dist})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	n := len(cities)
	matrix := make([][]int, n)
	data := make([]int, n*n)
	for i := range matrix {
		matrix[i] = data[i*n : (i+1)*n]
	}

	for _, edge := range edgeList {
		u, v, dist := edge[0], edge[1], edge[2]
		matrix[u][v] = dist
		matrix[v][u] = dist
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
	for i := range n {
		idx := (1<<i)*cols + i
		data[idx] = 0
		visited[idx] = true
	}

	// Main logic
	for mask := 1; mask < 1<<n; mask++ {
		for m := mask; m > 0; m &= m - 1 {
			last := bits.TrailingZeros(uint(m))
			prevMask := mask ^ (1 << last)
			if prevMask == 0 {
				continue
			}

			currIdx := mask*cols + last
			for pm := prevMask; pm > 0; pm &= pm - 1 {
				prev := bits.TrailingZeros(uint(pm))
				prevIdx := prevMask*cols + prev

				if !visited[prevIdx] {
					continue
				}

				newCost := data[prevIdx] + dist[prev][last]
				if !visited[currIdx] || compare(newCost, data[currIdx]) < 0 {
					data[currIdx] = newCost
					visited[currIdx] = true
				}
			}
		}
	}

	allVisited := rows - 1
	return slices.MinFunc(dp[allVisited], compare)
}

func main() {
	var partOne, partTwo int

	graph, err := ParseFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	partOne = Solve(graph, cmp.Compare[int])
	partTwo = Solve(graph, func(a, b int) int { return cmp.Compare(b, a) })
	fmt.Println(partOne)
	fmt.Println(partTwo)
}
