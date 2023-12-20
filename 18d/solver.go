package main

import (
	"fmt"
	"slices"
	"strings"
)

const (
	lava = 99
)

type Ope struct {
	dire byte
	dist int
}

type Color = string

func main() {
	ans1() // ans:
	// ans2() // ans:
}

func ans1() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	opes := make([]Ope, 0)
	colors := make([]Color, 0)
	digMap := make([][]int, 0)
	curPos := []int{0, 0}
	spots := make([][]int, 0)
	edges := []int{999, -999, -999, 999} // most U,D,R,L

	for ln := 0; scanner.Scan(); ln++ {
		line := scanner.Text()
		dire := line[0]
		dist := atoiEX(getSequenceNumber(line, 2))
		opes = append(opes, Ope{dire: dire, dist: dist})
		sharpIndex := strings.Index(line, "#")
		colors = append(colors, line[sharpIndex:len(line)-1])

		switch dire {
		case 'U':
			curPos[0] -= dist
			edges[0] = min(edges[0], curPos[0])
		case 'D':
			curPos[0] += dist
			edges[1] = max(edges[1], curPos[0])
		case 'R':
			curPos[1] += dist
			edges[2] = max(edges[2], curPos[1])
		case 'L':
			curPos[1] -= dist
			edges[3] = min(edges[3], curPos[1])
		}
		spots = append(spots, slices.Clone(curPos))
	}
	// fmt.Println(opes, colors)
	fmt.Println(spots)
	fmt.Println(edges)

	for i := 0; i < (edges[1] - edges[0]); i++ {
		digMap = append(digMap, make([]int, edges[2]-edges[3]))
	}
	// fmt.Println(digMap)

}
