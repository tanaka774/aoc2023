package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func ans2() {
	file, scanner := getScanner("./example.txt")
	// file, scanner := getScanner("./input.txt")
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

	conjs := make(map[string]int, 0) // "y,x":step
	for y, row := range island {
		for x, _ := range row {
			if isConjuction(island, y, x) {
				fmt.Println("conj", y, x)
				conjs[spf(y, x)] = 0
			}
		}
	}

	conjConnects := make(map[string]map[string]int)
	steps := make([]int, 0)
	temp := make(map[string]int)
	hike2(island, start[0], start[1], goal[0], goal[1], make(Set, 0), &steps, temp)
	conjConnects[spf(start[0], start[1])] = temp

	for conj := range conjs {
		y, x := splitNum(conj)
		temp := make(map[string]int)
		hike2(island, y, x, goal[0], goal[1], make(Set, 0), &steps, temp)
		conjConnects[spf(y, x)] = temp
	}

	maxPath := 0
	cycleConjunction(conjConnects, spf(start[0], start[1]), 0, &maxPath, make(map[string]bool), goal)
	fmt.Println("ans", maxPath)
}

func cycleConjunction(conjConnects map[string]map[string]int, conj string, paths int, maxPath *int, visited Set, goal []int) {
	if visited[conj] {
		return
	}
	visited[conj] = true

	y, x := splitNum(conj)
	if y == goal[0] && x == goal[1] {
		*maxPath = max(*maxPath, paths)
		return
	}

	for nconj, _ := range conjConnects[conj] {
		cycleConjunction(conjConnects, nconj, paths+conjConnects[conj][nconj], maxPath, copySet(visited), goal)
	}
}

func hike2(island []string, y, x, finishy, finishx int, went Set, steps *[]int, conjs map[string]int) {
	if y < 0 || y >= len(island) ||
		x >= len(island[y]) || x < 0 ||
		island[y][x] == '#' {
		return
	}

	if y == finishy && x == finishx {
		*steps = append(*steps, len(went))
		conjs[spf(y, x)] = max(conjs[spf(y, x)], len(went))
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
	if len(went) > 0 && isConjuction(island, y, x) {
		str := spf(y, x)
		conjs[str] = max(conjs[str], len(went))
		return
	}

	hike2(island, y-1, x, finishy, finishx, wentCopy, steps, conjs)
	hike2(island, y+1, x, finishy, finishx, wentCopy, steps, conjs)
	hike2(island, y, x-1, finishy, finishx, wentCopy, steps, conjs)
	hike2(island, y, x+1, finishy, finishx, wentCopy, steps, conjs)
}

func isConjuction(island []string, y, x int) bool {
	cnt := 0
	slopes := []byte{'^', 'v', '>', '<'}
	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, move := range moves {
		ny, nx := y+move[0], x+move[1]
		if ny < 0 || ny >= len(island) || nx < 0 || nx >= len(island[ny]) {
			continue
		}
		if slices.Contains(slopes, island[ny][nx]) {
			cnt++
		}
	}
	if cnt >= 3 {
		return true
	}
	return false
}

func atoi(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func spf(a, b int) string {
	return fmt.Sprintf("%d,%d", a, b)
}

func splitNum(str string) (a, b int) {
	ab := strings.Split(str, ",")
	a, b = atoi(ab[0]), atoi(ab[1])
	return
}

func copySet(s Set) (clone Set) {
	clone = make(map[string]bool)
	for k, v := range s {
		clone[k] = v
	}
	return
}
