package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// ans1() // ans:530495
	ans2() // ans:80253814
}

func ans1() {
	// readFile, fileScanner := getScanner("./example1_2.txt")
	readFile, fileScanner := getScanner("./input.txt")

	var schematic []string
	ans := 0 // sum of part numbers

	for fileScanner.Scan() {
		line := fileScanner.Text()
		schematic = append(schematic, line)
	}

	numberPlaces := map[string]struct{}{} // set structure

	for y, row := range schematic {
		for x, ch := range row {
			if !unicode.IsDigit(ch) && ch != '.' {
				places := findNumberPlaces(schematic, y, x)
				for _, place := range places {
					startIndex := getIndexBeforeSomethingToLeft(schematic[place[0]], place[1], unicode.IsDigit)
					numberPlaces[fmt.Sprintf("%d,%d", place[0], startIndex)] = struct{}{}
				}
			}
		}
	}
	// fmt.Println(numberPlaces)
	for place := range numberPlaces {
		yx := strings.Split(place, ",")
		y, _ := strconv.Atoi(yx[0])
		x, _ := strconv.Atoi(yx[1])
		// fmt.Println(y, x)
		num, _ := strconv.Atoi(getSequenceNumber(schematic[y], x))
		ans += num
	}

	fmt.Println("ans:", ans)

	readFile.Close()
}

func ans2() {
	// readFile, fileScanner := getScanner("./example2_1.txt")
	readFile, fileScanner := getScanner("./input.txt")

	var schematic []string
	ans := 0 // sum of part numbers

	for fileScanner.Scan() {
		line := fileScanner.Text()
		schematic = append(schematic, line)
	}

	for y, row := range schematic {
		for x, ch := range row {
			if ch == '*' {
				// fmt.Println("y, x at *", y, x)
				places := findNumberPlaces(schematic, y, x)
				// fmt.Println(places)
				if len(places) != 2 {
					continue
				}

				mul := 1
				for _, place := range places {
					startIndex := getIndexBeforeSomethingToLeft(schematic[place[0]], place[1], unicode.IsDigit)
					num, _ := strconv.Atoi(getSequenceNumber(schematic[place[0]], startIndex))
					mul *= num
				}
				ans += mul
			}
		}
	}

	fmt.Println("ans:", ans)

	readFile.Close()
}

func findNumberPlaces(strs []string, curY int, curX int) [][]int {
	var places [][]int
	if curX > 0 && unicode.IsDigit(rune(strs[curY][curX-1])) {
		places = append(places, []int{curY, curX - 1})
	}
	if curX > 0 && curY > 0 && unicode.IsDigit(rune(strs[curY-1][curX-1])) && !unicode.IsDigit(rune(strs[curY-1][curX])) {
		places = append(places, []int{curY - 1, curX - 1})
	}
	if curY > 0 && unicode.IsDigit(rune(strs[curY-1][curX])) {
		places = append(places, []int{curY - 1, curX})
	}
	if curY > 0 && curX < len(strs[curY-1])-1 && unicode.IsDigit(rune(strs[curY-1][curX+1])) && !unicode.IsDigit(rune(strs[curY-1][curX])) {
		places = append(places, []int{curY - 1, curX + 1})
	}
	if curX < len(strs[curY])-1 && unicode.IsDigit(rune(strs[curY][curX+1])) {
		places = append(places, []int{curY, curX + 1})
	}
	if curX < len(strs[curY+1])-1 && curY < len(strs)-1 && unicode.IsDigit(rune(strs[curY+1][curX+1])) && !unicode.IsDigit(rune(strs[curY+1][curX])) {
		places = append(places, []int{curY + 1, curX + 1})
	}
	if curY < len(strs)-1 && unicode.IsDigit(rune(strs[curY+1][curX])) {
		places = append(places, []int{curY + 1, curX})
	}
	if curX > 0 && curY < len(strs)-1 && unicode.IsDigit(rune(strs[curY+1][curX-1])) && !unicode.IsDigit(rune(strs[curY+1][curX])) {
		places = append(places, []int{curY + 1, curX - 1})
	}
	return places
}
