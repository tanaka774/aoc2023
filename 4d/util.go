package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func getScanner(fileName string) (*os.File, *bufio.Scanner) {
	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return readFile, fileScanner
}

func getSequenceNumber(str string, startIndex int) string {
	num := ""
	for {
		if startIndex > len(str)-1 {
			break
		}

		if unicode.IsDigit(rune(str[startIndex])) {
			num += fmt.Sprint(getNumber(str[startIndex]))
		} else {
			break
		}
		startIndex++
	}

	return num
}

func getNumber(b byte) int {
	return int(b) - '0'
}

func findIndexes(str string, searchWord string) []int {
	var indexes []int

	startIndex := 0
	for {
		index := strings.Index(str[startIndex:], searchWord)
		if index == -1 {
			break
		}

		indexes = append(indexes, startIndex+index)
		startIndex += index + 1
	}

	return indexes
}

func getIndexBeforeSomethingToLeft(str string, startIndex int, callback func(r rune) bool) int {
	index := startIndex
	for {
		if index == 0 || !callback(rune(str[index-1])) {
			break
		}
		index--
	}
	return index
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
