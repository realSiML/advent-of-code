package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	size  = 100
	steps = 100
)

type Grid [size][size]bool

func ParseFile(path string) Grid {
	var lights Grid

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	i := 0
	for sc.Scan() {
		for j, c := range sc.Bytes() {
			if c == '#' {
				lights[i][j] = true
			}
		}
		i++
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return lights
}

func Tick(lights *Grid) {
	snapshot := *lights

	countNeighbors := func(i, j int) int {
		count := 0
		for di := -1; di <= 1; di++ {
			for dj := -1; dj <= 1; dj++ {
				if (0 <= i+di && i+di < size) &&
					(0 <= j+dj && j+dj < size) &&
					(i+di != i || j+dj != j) &&
					snapshot[i+di][j+dj] {
					count++
				}
			}
		}
		return count
	}

	for i := range lights {
		for j := range lights[i] {
			neighbors := countNeighbors(i, j)
			if snapshot[i][j] && !(neighbors == 2 || neighbors == 3) {
				(*lights)[i][j] = false
			} else if !snapshot[i][j] && neighbors == 3 {
				(*lights)[i][j] = true
			}
		}
	}
}

func CountLights(lights Grid) int {
	count := 0
	for i := range lights {
		for j := range lights[i] {
			if lights[i][j] {
				count++
			}
		}
	}
	return count
}

func LightCorners(lights *Grid) {
	(*lights)[0][0] = true
	(*lights)[0][size-1] = true
	(*lights)[size-1][0] = true
	(*lights)[size-1][size-1] = true
}

func main() {
	lights := ParseFile("./input.txt")
	for range steps {
		Tick(&lights)
	}
	partOne := CountLights(lights)
	fmt.Println(partOne)

	lights = ParseFile("./input.txt")
	LightCorners(&lights)
	for range steps {
		Tick(&lights)
		LightCorners(&lights)
	}
	partTwo := CountLights(lights)
	fmt.Println(partTwo)
}
