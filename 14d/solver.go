package main

import (
	"fmt"
	"slices"
)

const (
	NORTH = "NORTH"
	SOUTH = "SOUTH"
	WEST  = "WEST"
	EAST  = "EAST"
)

func main() {
	// ans1() // ans:108857
	ans2() // ans:
}

func ans1() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	platform := make([]string, 0)
	for ln := 0; scanner.Scan(); ln++ {
		line := scanner.Text()
		platform = append(platform, line)
	}
	// printArray[string](platform, "pf")

	tilted := slices.Clone(platform)

	for y, row := range platform {
		rockPoss := findIndexes(row, "O")
		for _, rockPos := range rockPoss {
			tilted[y] = tilted[y][:rockPos] + "." + tilted[y][rockPos+1:]
			for yi := 1; yi < len(platform); yi++ {
				// fmt.Println(y, yi, rockPos)
				if y-yi < 0 || tilted[y-yi][rockPos] == 'O' || tilted[y-yi][rockPos] == '#' {
					// fmt.Println("stop here", y, yi, rockPos)
					tilted[y-yi+1] = tilted[y-yi+1][:rockPos] + "O" + tilted[y-yi+1][rockPos+1:]
					// printArray[string](tilted, "til")
					break
				}
			}
		}
	}

	ans := 0
	printArray[string](tilted, "til")
	for y, row := range tilted {
		rockPoss := findIndexes(row, "O")
		ans += len(rockPoss) * (len(tilted) - y)

	}
	fmt.Println("ans", ans)
}

func ans2() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	platform := make([]string, 0)
	for ln := 0; scanner.Scan(); ln++ {
		line := scanner.Text()
		platform = append(platform, line)
	}
	// printArray[string](platform, "pf")

	tilted := slices.Clone(platform)

	// loopStart := 2  // example
	// loopIndice := 7 // example
	// loopStart := 104 // input
	// loopIndice := 13 // input
	// for i := 0; i < loopStart+((1000000000-loopStart)%loopIndice); i++ {
	for i := 0; i < 40; i++ {
		tilted = tiltNorth(tilted)
		// fmt.Println("N cnt:", countRocks(tilted))
		printArray(tilted, "N til")
		tilted = tiltWest(tilted)
		// fmt.Println("W cnt:", countRocks(tilted))
		printArray(tilted, "W til")
		tilted = tiltSouth(tilted)
		// fmt.Println("S cnt:", countRocks(tilted))
		printArray(tilted, "S til")
		tilted = tiltEast(tilted)
		// fmt.Println("E cnt:", countRocks(tilted))
		printArray(tilted, "E til")
		ans := 0
		// printArray[string](tilted, "til")
		for y, row := range tilted {
			rockPoss := findIndexes(row, "O")
			ans += len(rockPoss) * (len(tilted) - y)

		}
		fmt.Println("i:ans", i, ans)
		// from i:104 per 14items, ans:95274~95273
	}

	// ans := 0
	// printArray[string](tilted, "til")
	// for y, row := range tilted {
	// 	rockPoss := findIndexes(row, "O")
	// 	ans += len(rockPoss) * (len(tilted) - y)
	//
	// }
	// fmt.Println("ans", ans)
}

func countRocks(tilted []string) (count int) {
	for _, row := range tilted {
		rockPoss := findIndexes(row, "O")
		count += len(rockPoss)
	}
	return
}

func tiltNorth(platform []string) []string {
	tilted := slices.Clone(platform)
	for y, row := range tilted {
		rockPoss := findIndexes(row, "O")
		for _, rockPos := range rockPoss {
			tilted[y] = tilted[y][:rockPos] + "." + tilted[y][rockPos+1:]
			for yi := 1; yi < len(tilted); yi++ {
				if y-yi < 0 || tilted[y-yi][rockPos] == 'O' || tilted[y-yi][rockPos] == '#' {
					tilted[y-yi+1] = tilted[y-yi+1][:rockPos] + "O" + tilted[y-yi+1][rockPos+1:]
					break
				}
			}
		}
	}
	return tilted
}

func tiltSouth(platform []string) []string {
	tilted := slices.Clone(platform)
	for y := len(tilted) - 1; y >= 0; y-- {
		rockPoss := findIndexes(tilted[y], "O")
		for _, rockPos := range rockPoss {
			tilted[y] = tilted[y][:rockPos] + "." + tilted[y][rockPos+1:]
			for yi := 1; yi < len(tilted); yi++ {
				if y+yi > len(tilted)-1 || tilted[y+yi][rockPos] == 'O' || tilted[y+yi][rockPos] == '#' {
					tilted[y+yi-1] = tilted[y+yi-1][:rockPos] + "O" + tilted[y+yi-1][rockPos+1:]
					break
				}
			}
		}
	}
	return tilted
}

func tiltWest(platform []string) []string {
	tilted := slices.Clone(platform)
	for y, row := range platform {
		for x, ch := range row {
			if ch != 'O' {
				continue
			}
			tilted[y] = tilted[y][:x] + "." + tilted[y][x+1:]
			for xi := 1; xi <= len(tilted[y]); xi++ {
				// if y == 7 {
				// 	fmt.Println("missing", y, x, xi)
				// }
				if x-xi < 0 || tilted[y][x-xi] == 'O' || tilted[y][x-xi] == '#' {
					tilted[y] = tilted[y][:x-xi+1] + "O" + tilted[y][x-xi+2:]
					break
				}
			}
		}
	}
	return tilted
}

func tiltEast(platform []string) []string {
	tilted := slices.Clone(platform)
	for y, row := range platform {
		for x := len(row) - 1; x >= 0; x-- {
			if platform[y][x] != 'O' {
				continue
			}
			tilted[y] = tilted[y][:x] + "." + tilted[y][x+1:]
			for xi := 1; xi < len(tilted[y]); xi++ {
				if x+xi > len(tilted[y])-1 || tilted[y][x+xi] == 'O' || tilted[y][x+xi] == '#' {
					tilted[y] = tilted[y][:x+xi-1] + "O" + tilted[y][x+xi:]
					break
				}
			}
		}
	}
	return tilted
}
