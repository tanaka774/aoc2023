package main

import (
	"fmt"
)

const (
	LEFT = iota
	RIGHT
	UP
	DOWN
	VERBOTH
	HORBOTH
)

type set = map[string]bool

func main() {
	ans1() // ans:7884
	// ans2() // ans:
}

func ans1() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	tiles := make([]string, 0)
	for ln := 0; scanner.Scan(); ln++ {
		line := scanner.Text()
		tiles = append(tiles, line)
	}

	went := make(set, 0) // "y,x,dire"
	beam(tiles, 0, 0, RIGHT, went)

	wentForCount := make(set, 0) //"y,x"
	for key := range went {
		yx := key[:len(key)-2]
		wentForCount[yx] = true
	}
	fmt.Println(len(wentForCount))
}

func beam(tiles []string, y int, x int, dire int, went set) {
	if y < 0 || y > len(tiles)-1 || x < 0 || x > len(tiles[y])-1 {
		return
	}
	// fmt.Println("y x dire sym:", y, x, dire, string(tiles[y][x]))

	str := fmt.Sprintf("%d,%d,%d", y, x, dire)
	if went[str] {
		return
	}
	went[str] = true

	ndires := map[byte][]int{
		'.':  {LEFT, RIGHT, UP, DOWN},
		'\\': {UP, DOWN, LEFT, RIGHT},
		'/':  {DOWN, UP, RIGHT, LEFT},
		'|':  {VERBOTH, VERBOTH, UP, DOWN},
		'-':  {LEFT, RIGHT, HORBOTH, HORBOTH},
	}

	ndire := ndires[tiles[y][x]][dire]
	moves := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // l,r,u,d

	switch ndire {
	case VERBOTH:
		beam(tiles, y-1, x, UP, went)
		beam(tiles, y+1, x, DOWN, went)
	case HORBOTH:
		beam(tiles, y, x-1, LEFT, went)
		beam(tiles, y, x+1, RIGHT, went)
	default:
		ny := y + moves[ndire][0]
		nx := x + moves[ndire][1]
		beam(tiles, ny, nx, ndire, went)
	}
}
