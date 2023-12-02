package main

import (
	"fmt"
	"strconv"
)

func main() {
	ans1() // ans:2632
	ans2() // ans:69269
}

func ans1() {
	readFile, fileScanner := getScanner("./input.txt")

	colors := [3]string{"red", "green", "blue"}
	maxs := [3]int{12, 13, 14}
	ans1 := 0 // sum of games

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)

		currentGame, _ := strconv.Atoi(getSequenceNumber(line, 5))
		isSmallBag := true

		for i, color := range colors {
			indexes := findIndexes(line, color)
			for _, index := range indexes {
				numIndex := getIndexBeforeSpaceToLeft(line, index-2)
				colorNum, _ := strconv.Atoi(getSequenceNumber(line, numIndex))

				if colorNum > maxs[i] {
					isSmallBag = false
					break
				}

			}
			if !isSmallBag {
				break
			}
		}

		if isSmallBag {
			ans1 += currentGame
		}
	}

	fmt.Printf("ans1: %d\n", ans1)

	readFile.Close()
}

func ans2() {
	readFile, fileScanner := getScanner("./input.txt")

	colors := [3]string{"red", "green", "blue"}
	ans2 := 0 // sum of multiplied color numbers

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)

		var maxs [3]int
		for i, color := range colors {
			indexes := findIndexes(line, color)
			for _, index := range indexes {
				numIndex := getIndexBeforeSpaceToLeft(line, index-2)
				colorNum, _ := strconv.Atoi(getSequenceNumber(line, numIndex))
				maxs[i] = max(maxs[i], colorNum)
			}
		}

		// fmt.Println("maxs:", maxs)
		mul := 1
		for _, m := range maxs {
			mul *= m
		}
		// fmt.Println("mul:", mul)

		ans2 += mul
	}

	fmt.Printf("ans2: %d\n", ans2)

	readFile.Close()
}
