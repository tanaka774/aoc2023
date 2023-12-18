package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

func main() {
	ans1() // ans:
	// ans2() // ans:
}

const (
	LEFT int = iota
	UP
	RIGHT
	DOWN
	UNINIT
)

var (
	direList = []int{LEFT, UP, RIGHT, DOWN} // 0,1,2,3
	moves    = [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
)

type Pos struct {
	y, x int
}

type Lava struct {
	pos      Pos
	dire     int // L,R,U,D
	count    int // at most 3 on the same direction
	heatLoss int
}

type Blocks = []string

// type set = map[string]bool

// func (went set) alreadyHas(y, x int) bool {
// 	str := fmt.Sprintf("%d,%d", y, x)
// 	return went[str]
// }

// type cityHandler interface {
// 	getHeatLoss(pos Pos)
// }

func (pos Pos) HeatLoss(blocks Blocks) byte {
	return blocks[pos.y][pos.x]
}

func ans1() {
	// file, scanner := getScanner("./ex2.txt")
	file, scanner := getScanner("./example.txt")
	// file, scanner := getScanner("./input.txt")
	defer file.Close()

	blocks := make(Blocks, 0)
	for ln := 0; scanner.Scan(); ln++ {
		line := scanner.Text()
		blocks = append(blocks, line)
	}

	// lava := Lava{Pos{0, 0}, UNINIT, 0, 0}
	// costStore := make([]int, 0)
	costMap := make(map[string]int, 0)
	for y := 0; y < len(blocks); y++ {
		for x := 0; x < len(blocks[y]); x++ {
			costMap[fmt.Sprintf("%d,%d", y, x)] = math.MaxUint32
		}
	}
	directing(blocks, 0, 0, UNINIT, 0, make([][]int, 0), costMap)
}

// func directing(blocks Blocks, y, x, preDire, count, totalHeatLoss int, temps [][]int, costMap map[string]int) {
func directing(blocks Blocks, y, x, preDire, count int, temps [][]int, costMap map[string]int) {
	str := fmt.Sprintf("%d,%d", y, x)
	if len(temps) > 0 && len(temps[0]) > 0 {
		// fmt.Println("st", costStore[len(costStore)-1], costMap[str])
		if costMap[str] <= temps[0][len(temps[0])-1] {
			// compare heatloss
			// fmt.Println("visited", y, x, count, temps[0][4], temps[0][len(temps[0])-1], costMap[str])
			return
		} else {
			fmt.Println("update", y, x, count, temps[0][4], temps[0][len(temps[0])-1], costMap[str])
			costMap[str] = temps[0][len(temps[0])-1]
		}
	}

	if y == len(blocks)-1 && x == len(blocks[y])-1 {
		fmt.Println("finish?", temps[0][4])
		return
	}
	dires := make([]int, 0)
	if preDire == UNINIT {
		dires = slices.Clone(direList)
	} else if count >= 3 {
		dires = append(dires, direList[(preDire+1)%4], direList[(preDire+3)%4])
	} else {
		dires = append(dires, preDire, direList[(preDire+1)%4], direList[(preDire+3)%4])
	}
	// fmt.Println("dire", preDire, dires)

	// temps := make([][]int, 0) // ny, nx, dire, heatLoss, cost
	for _, dire := range dires {
		ny := y + moves[dire][0]
		nx := x + moves[dire][1]
		if ny < 0 || nx < 0 || ny >= len(blocks) || nx >= len(blocks[ny]) {
			continue //not return
		}
		if dire != preDire {
			count = 0
		}
		count++

		sdist := ny - 0 + nx - 0
		gdist := len(blocks) - 1 - ny + len(blocks) - 1 - nx
		heatLoss := getNumber(blocks[ny][nx])
		totalHeatLoss := 0
		if len(temps) > 1 && len(temps[0]) > 0 {
			totalHeatLoss += temps[0][4]
		}
		totalHeatLoss += heatLoss
		rate := 4 // dist:heatLoss
		cost := totalHeatLoss + (sdist+gdist)*rate
		// fmt.Println("cost", heatLoss, sdist, gdist, rate, cost)

		temps = append(temps, []int{ny, nx, dire, count, totalHeatLoss, cost})
	}

	// sort order by cost asc
	sort.Slice(temps, func(a, b int) bool {
		return temps[a][len(temps[a])-1] < temps[b][len(temps[b])-1]
	})

	// fmt.Println("deb", temps)

	// count++
	for _, temp := range temps {
		ny, nx, dire, count := temp[0], temp[1], temp[2], temp[3]
		// if dire != preDire {
		// 	count = 0
		// }
		// costStore = append(costStore, cost)
		// totalHeatLoss += heatLoss

		directing(blocks, ny, nx, dire, count, temps, costMap)

		temps = temps[1:]
		// fmt.Println("out", temp)

		// costStore = costStore[:len(costStore)-1]
		// totalHeatLoss -= getNumber(blocks[y][x])
	}
}
