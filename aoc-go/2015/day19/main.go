package main

import (
	"bufio"
	"fmt"
	"iter"
	"log"
	"os"
	"regexp"
	"strings"
)

type rule struct {
	from string
	to   string
}

func ParseFile(path string) ([]rule, string) {
	rules, molecule := []rule{}, ""

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		parts := strings.Fields(line)
		if len(parts) == 3 {
			from, to := parts[0], parts[2]
			rules = append(rules, rule{from, to})
		} else if line != "" {
			molecule = line
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return rules, molecule
}

func GetMoleculeVariants(s, old, new string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for i := 0; i+len(old) <= len(s); i++ {
			if s[i:i+len(old)] == old {
				r := s[:i] + new + s[i+len(old):]
				if !yield(r) {
					return
				}
			}
		}
	}
}

func PartOne(rules []rule, molecule string) int {
	m := map[string]struct{}{}
	for _, r := range rules {
		for replaced := range GetMoleculeVariants(molecule, r.from, r.to) {
			m[replaced] = struct{}{}
		}
	}
	return len(m)
}

func countAtoms(s string) int {
	re := regexp.MustCompile(`[A-Z][a-z]?`)
	return len(re.FindAllString(s, -1))
}

func PartTwo(molecule string) int {
	atoms := countAtoms(molecule)
	rn := strings.Count(molecule, "Rn")
	ar := strings.Count(molecule, "Ar")
	y := strings.Count(molecule, "Y")

	return atoms - rn - ar - 2*y - 1
}

func main() {
	rules, molecule := ParseFile("./input.txt")

	partOne := PartOne(rules, molecule)
	fmt.Println(partOne)
	partTwo := PartTwo(molecule)
	fmt.Println(partTwo)
}
