package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	ans1() // ans:55971
	// ans2() // ans:54719
}

func ans1() {
	ReadFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(ReadFile)
	fileScanner.Split(bufio.ScanLines)

	ans := 0 // sum of "first + last"

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)

		first := ""
		last := ""
		zero := 48
		for i := range line {
			if unicode.IsDigit(rune(line[i])) && first == "" {
				first = fmt.Sprint(int(line[i]) - zero)
				// fmt.Printf("f:%s ", first)
			}
			if unicode.IsDigit(rune(line[len(line)-1-i])) && last == "" {
				last = fmt.Sprint(int(line[len(line)-1-i]) - zero)
				// fmt.Printf("l:%s ", last)
			}
			if first != "" && last != "" {
				break
			}
		}

		sum, _ := strconv.Atoi(first + last)
		ans += sum
		// fmt.Println("")
	}

	fmt.Println("ans1:", ans)

	ReadFile.Close()
}

func ans2() {
	ReadFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(ReadFile)
	fileScanner.Split(bufio.ScanLines)

	digits := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	ans := 0 // sum of "first + last"

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)

		first := ""
		last := ""
		zero := 48
		for lineIndex := range line {
			if first == "" {
				if unicode.IsDigit(rune(line[lineIndex])) {
					first = fmt.Sprint(int(line[lineIndex]) - zero)
				} else {
					for digitIndex, digit := range digits {
						if isThisWord(line, lineIndex, digit) {
							first = strconv.Itoa(digitIndex + 1)
						}
					}
				}
			}

			if last == "" {
				if unicode.IsDigit(rune(line[len(line)-1-lineIndex])) {
					last = fmt.Sprint(int(line[len(line)-1-lineIndex]) - zero)
				} else {
					for digitIndex, digit := range digits {
						if isThisWord(line, len(line)-1-lineIndex, digit) {
							last = strconv.Itoa(digitIndex + 1)
						}
					}
				}
			}

			if first != "" && last != "" {
				// fmt.Printf("f:%s ", first)
				// fmt.Printf("l:%s \n", last)
				break
			}
		}

		sum, _ := strconv.Atoi(first + last)
		ans += sum
	}

	fmt.Println("ans2:", ans)

	ReadFile.Close()
}

func isThisWord(str string, index int, searchWord string) bool {
	hereItIs := true
	for i, searchChar := range searchWord {
		if index+i >= len(str) {
			return false
		}
		if searchChar != rune(str[index+i]) {
			hereItIs = false
		}
	}
	return hereItIs
}
