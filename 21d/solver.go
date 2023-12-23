package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	y int
	x int
}
type Garden = []string
type Set = map[string]bool

const (
	MAX_STEP = 6
	UP       = "UP"
	DOWN     = "DOWN"
	LEFT     = "LEFT"
	RIGHT    = "RIGHT"
	INIT     = "INIT"
)

func main() {
	// f, err := os.Open("example.txt")
	f, err := os.Open("input.txt") // ans1:3594
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// scanner.Split(bufio.ScanWords) // get each char

	garden := make(Garden, 0)
	var start Pos

	for ln := 0; scanner.Scan(); ln++ {
		// fmt.Println(scanner.Text())
		line := scanner.Text()
		garden = append(garden, line)
		si := strings.Index(line, "S")
		if si != -1 {
			start = Pos{y: ln, x: si}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	printSlice(garden)
	fmt.Println(start)

	possibles := make(Set, 0)
	possibles[fmt.Sprintf("%d,%d", start.y, start.x)] = true

	// cur := slices.Clone(start)
	for i := 0; i < 64; i++ {
		for _, pos := range getMapKeys(possibles) {
			yx := strings.Split(pos, ",")
			y, x := atoi(yx[0]), atoi(yx[1])
			planting(garden, y+1, x, UP, possibles)
			planting(garden, y-1, x, DOWN, possibles)
			planting(garden, y, x+1, LEFT, possibles)
			planting(garden, y, x-1, RIGHT, possibles)
		}

	}
	fmt.Println(possibles)
	fmt.Println(len(possibles))
}

// func planting(garden Garden, y, x int, prev string, possibles Set, steps int) {
func planting(garden Garden, y, x int, prev string, possibles Set) {
	if
	// steps > MAX_STEP ||
	y < 0 || y >= len(garden) ||
		x >= len(garden[y]) || x < 0 ||
		garden[y][x] == '#' {
		return
	}
	// fmt.Println("deb:", y, x, prev, possibles)
	str := fmt.Sprintf("%d,%d", y, x)
	// // this stops later branch
	// if possibles[str] {
	// 	return
	// }
	possibles[str] = true

	switch prev {
	case UP:
		// possibles[fmt.Sprintf("%d,%d", y-1, x)] = false
		delete(possibles, fmt.Sprintf("%d,%d", y-1, x))
	case DOWN:
		// possibles[fmt.Sprintf("%d,%d", y+1, x)] = false
		delete(possibles, fmt.Sprintf("%d,%d", y+1, x))
	case LEFT:
		// possibles[fmt.Sprintf("%d,%d", y, x-1)] = false
		delete(possibles, fmt.Sprintf("%d,%d", y, x-1))
	case RIGHT:
		// possibles[fmt.Sprintf("%d,%d", y, x+1)] = false
		delete(possibles, fmt.Sprintf("%d,%d", y, x+1))
	}

	// planting(garden, y+1, x, UP, possibles, steps+1)
	// planting(garden, y-1, x, DOWN, possibles, steps+1)
	// planting(garden, y, x+1, LEFT, possibles, steps+1)
	// planting(garden, y, x-1, RIGHT, possibles, steps+1)
}

func atoi(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func getMapKeys[T comparable, R any](m map[T]R) (keys []T) {
	// keys2 := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return
}

func printSlice[T any](sl []T) {
	for _, v := range sl {
		fmt.Println(v)
	}
}
