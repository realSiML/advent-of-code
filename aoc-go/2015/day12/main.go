package main

import (
	"encoding/json/jsontext"
	"fmt"
	"io"
	"log"
	"os"
)

func PartOne(file *os.File) int64 {
	var total int64

	dec := jsontext.NewDecoder(file)
	for {
		tok, err := dec.ReadToken()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if tok.Kind() == '0' {
			if num, err := tok.Int(); err == nil {
				total += num
			}
		}
	}

	return total
}

func PartTwo(file *os.File) int64 {
	sums := []int64{0}
	dec := jsontext.NewDecoder(file)

	for {
		tok, err := dec.ReadToken()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		switch tok.Kind() {
		case '{':
			sums = append(sums, 0)
		case '}':
			last := sums[len(sums)-1]
			sums = sums[:len(sums)-1]
			sums[len(sums)-1] += last
		case '0':
			if num, err := tok.Int(); err == nil {
				sums[len(sums)-1] += num
			}
		case '"':
			depth := dec.StackDepth()
			if depth > 0 {
				kind, index := dec.StackIndex(depth)
				if kind == '{' && index%2 == 0 && tok.String() == "red" {
					sums = sums[:len(sums)-1]
					skipCurrentContainer(dec, depth)
				}
			}
		}
	}

	return sums[0]
}

func skipCurrentContainer(dec *jsontext.Decoder, targetDepth int) {
	for dec.StackDepth() >= targetDepth {
		if _, err := dec.ReadToken(); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(PartOne(file))
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(PartTwo(file))
}
