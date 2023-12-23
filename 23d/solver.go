package main

import (
	"fmt"
	"slices"
	"strings"
)

const (
	UP int = iota
	DOWN
	LEFT
	RIGHT
)

type Set = map[string]bool

func main() {
	// ans1() // ans:1930
	ans2() // ans:6230
}

func ans1() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	island := make([]string, 0)
	start := []int{0, 0}
	goal := []int{len(island) - 1, 0}
	for ln := 0; scanner.Scan(); ln++ {
		line := scanner.Text()
		island = append(island, line)
		if ln == 0 {
			p := strings.Index(line, ".")
			start[1] = p
		}
		if ln == len(island)-1 {
			p := strings.Index(line, ".")
			goal[0], goal[1] = ln, p
		}
	}
	fmt.Println(start)
	fmt.Println(goal)
	steps := make([]int, 0)
	// hike(island, start[0], start[1], make(Set, 0), &steps)
	hike(island, goal[0], goal[1], make(Set, 0), &steps)
	fmt.Println("ans", slices.Max(steps))
}

func hike(island []string, y, x int, went Set, steps *[]int) {
	if y < 0 || y >= len(island) ||
		x >= len(island[y]) || x < 0 ||
		island[y][x] == '#' {
		return
	}

	// if y == len(island)-1 {
	if y == 0 {
		fmt.Println("goal", len(went))
		*steps = append(*steps, len(went))
		return
	}

	wentCopy := make(Set, 0)
	for k, v := range went {
		wentCopy[k] = v
	}

	str := fmt.Sprintf("%d,%d", y, x)
	if wentCopy[str] {
		return
	}
	wentCopy[str] = true

	switch island[y][x] {
	// // part1
	// case '^':
	// 	hike(island, y-1, x, wentCopy, steps)
	// case 'v':
	// 	hike(island, y+1, x, wentCopy, steps)
	// case '<':
	// 	hike(island, y, x-1, wentCopy, steps)
	// case '>':
	// 	hike(island, y, x+1, wentCopy, steps)
	default: // .
		hike(island, y-1, x, wentCopy, steps)
		hike(island, y+1, x, wentCopy, steps)
		hike(island, y, x-1, wentCopy, steps)
		hike(island, y, x+1, wentCopy, steps)
	}

	//remove went?
	// for k := range wentCopy {
	// 	delete(wentCopy, k)
	// }

}
