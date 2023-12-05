package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	ans1() // ans:21138
	// ans2() // ans:7185540
}

func ans1() {
	// readFile, fileScanner := getScanner("./example.txt")
	readFile, fileScanner := getScanner("./input.txt")

	ans := 0 // sum of match points
	for fileScanner.Scan() {
		line := fileScanner.Text()

		winNums := make([]int, 0)
		myNums := make([]int, 0)
		isMyTurn := false // become true after passing '|'

		curIndex := strings.Index(line, ":") + 1
		for {
			for {
				if string(line[curIndex]) == "|" {
					isMyTurn = true
					curIndex++
				}
				if string(line[curIndex]) == " " {
					curIndex++
				} else {
					break
				}
			}

			numStr := getSequenceNumber(line, curIndex)

			num, _ := strconv.Atoi(numStr)
			if isMyTurn {
				myNums = append(myNums, num)
			} else {
				winNums = append(winNums, num)
			}
			curIndex += len(numStr)

			// TODO want to check just once
			if curIndex >= len(line)-1 {
				break
			}
		}

		matchCount := 0
		for _, myNum := range myNums {
			for _, winNum := range winNums {
				if myNum == winNum {
					matchCount++
				}
			}
		}
		if matchCount >= 1 {
			ans += int(math.Pow(2.0, float64(matchCount-1)))
		}
	}

	fmt.Println("ans:", ans)

	readFile.Close()
}

func ans2() {
	// readFile, fileScanner := getScanner("./example.txt")
	readFile, fileScanner := getScanner("./input.txt")

	ans := 0 // sum of all cards
	copyCards := []int{0}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		winNums := make([]int, 0)
		myNums := make([]int, 0)
		isMyTurn := false // become true after passing '|'

		curIndex := strings.Index(line, ":") + 1
		for {
			for {
				if string(line[curIndex]) == "|" {
					isMyTurn = true
					curIndex++
				}
				if string(line[curIndex]) == " " {
					curIndex++
				} else {
					break
				}
			}

			numStr := getSequenceNumber(line, curIndex)

			num, _ := strconv.Atoi(numStr)
			if isMyTurn {
				myNums = append(myNums, num)
			} else {
				winNums = append(winNums, num)
			}
			curIndex += len(numStr)

			if curIndex >= len(line)-1 {
				break
			}
		}

		matchCount := 0
		for _, myNum := range myNums {
			for _, winNum := range winNums {
				if myNum == winNum {
					matchCount++
				}
			}
		}
		var allCards int

		if len(copyCards) >= 1 {
			allCards = 1 + copyCards[0]
			copyCards = copyCards[1:]
		} else {
			allCards = 1
			temp := []int{0}
			copyCards = temp[:]
		}

		copyCards = handleCopyCards(copyCards, matchCount, allCards)

		ans += allCards
	}

	fmt.Println("ans:", ans)

	readFile.Close()
}

func handleCopyCards(copyCards []int, matchCount int, allCards int) []int {
	if matchCount == 0 {
		return copyCards[:]
	}
	nextAddedCards := make([]int, matchCount)
	for i := range nextAddedCards {
		nextAddedCards[i] = allCards
	}

	var longer []int
	var shorter []int

	if len(copyCards) >= len(nextAddedCards) {
		longer = copyCards[:]
		shorter = nextAddedCards[:]
	} else {
		longer = nextAddedCards[:]
		shorter = copyCards[:]
	}

	res := longer[:]

	for i, sh := range shorter {
		res[i] = longer[i] + sh
	}

	return res
}
