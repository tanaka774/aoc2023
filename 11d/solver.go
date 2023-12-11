package main

import (
	"fmt"
	"math"
)

const (
	// EXPANSION = 1 // ans1
	// EXPANSION = 10 - 1 // ans2 ex-1
	// EXPANSION = 100 - 1 // ans2 ex-2
	EXPANSION = 1000000 - 1 // ans2
)

func main() {
	ans() // ans1:9591768 ans2:746962097860
}

func ans() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	rowsNoStar := make([]int, 0)
	colsNoStar := make([]int, 0)
	stars := make([][]int, 0)

	for ln := 0; scanner.Scan(); ln++ {
		line := scanner.Text()

		if ln == 0 {
			for n := 0; n < len(line); n++ {
				colsNoStar = append(colsNoStar, n)
			}
		}

		starXPoss := findIndexes(line, "#")
		if len(starXPoss) == 0 {
			rowsNoStar = append(rowsNoStar, ln)
		} else {
			for _, x := range starXPoss {
				colsNoStar = deleteElemWithElem(colsNoStar, x)
				stars = append(stars, []int{x, ln})
			}
		}

	}

	for i := range stars {
		for ri := range rowsNoStar {
			if stars[i][1] > rowsNoStar[len(rowsNoStar)-1-ri] {
				stars[i][1] += EXPANSION
			}
		}
		for ci := range colsNoStar {
			if stars[i][0] > colsNoStar[len(colsNoStar)-1-ci] {
				stars[i][0] += EXPANSION
			}
		}
	}

	ans := 0
	for i := range stars {
		for j := i + 1; j < len(stars); j++ {
			ans += int(math.Abs(float64(stars[i][0] - stars[j][0])))
			ans += int(math.Abs(float64(stars[i][1] - stars[j][1])))
		}
	}
	fmt.Println("ans", ans)
}
