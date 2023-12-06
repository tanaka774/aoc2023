package main

import (
	"fmt"
	"slices"
	"strings"
	"unicode"
)

const (
	seedLen = 7
)

func main() {
	// ans1() // ans:322500873
	ans2() // ans:108956227
}

func ans1() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	seeds := make([]int, 0)
	var seedsConverted []bool
	var currents []int
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "seeds:") {
			seeds = getNums(line, seedLen)
			currents = seeds[:]
			seedsConverted = make([]bool, len(seeds))
		}

		if line == "" {
			for i := range seedsConverted {
				seedsConverted[i] = false
			}
			continue
		}
		if !unicode.IsDigit(rune(line[0])) {
			continue
		}

		change := getNums(line, 0)

		for i, cur := range currents {
			if !seedsConverted[i] && cur >= change[1] && cur <= (change[1]+change[2]) {
				currents[i] = change[0] + (cur - change[1])
				seedsConverted[i] = true
			}
		}
	}

	m := slices.Min(currents)
	fmt.Println("ans:", m)
}

func ans2() {
	file, scanner := getScanner("./example.txt")
	// file, scanner := getScanner("./input.txt")
	defer file.Close()

	seeds := make([]int, 0)
	var seedsConverted []bool
	var currents []int
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "seeds:") {
			seeds = getNums(line, seedLen)
			var actualSeeds []int
			for i, seed := range seeds {
				if i%2 == 1 {
					continue
				}

				for a := 0; a < seeds[i+1]; a++ {
					actualSeeds = append(actualSeeds, seed+a)
				}
			}

			currents = actualSeeds[:]
			seedsConverted = make([]bool, len(actualSeeds))
		}

		if line == "" {
			for i := range seedsConverted {
				seedsConverted[i] = false
			}
			continue
		}
		if !unicode.IsDigit(rune(line[0])) {
			continue
		}

		change := getNums(line, 0)

		for i, cur := range currents {
			if !seedsConverted[i] && cur >= change[1] && cur <= (change[1]+change[2]) {
				currents[i] = change[0] + (cur - change[1])
				seedsConverted[i] = true
			}
		}
	}

	m := slices.Min(currents)
	fmt.Println("ans:", m)
}
