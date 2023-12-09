package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"slices"
	"strconv"
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

/**
* use if the space is just one between numbers
 */
func getNums(line string, firstIndex int) []int {
	nums := make([]int, 0)
	curIndex := firstIndex

	for {
		numStr := getSequenceNumber(line, curIndex)
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)

		curIndex += len(numStr) + 1
		if curIndex >= len(line) {
			break
		}
	}

	return nums
}

func atoiEX(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func sortString(str string) string {
	runes := []rune(str)
	slices.Sort(runes)
	return string(runes)
}

func every[T any](arr []T, fn func(T) bool) bool {
	for _, element := range arr {
		if !fn(element) {
			return false
		}
	}
	return true
}

// func calcGCD(nums ...int64) *big.Int {
func calcGCD(nums []int64) *big.Int {
	if len(nums) == 0 {
		return big.NewInt(0)
	}
	gcd := big.NewInt(nums[0])

	for _, num := range nums[1:] {
		gcd = gcd.GCD(nil, nil, gcd, big.NewInt(num))
	}

	return gcd
}

// type temp interface {
//   []int64 | ...int64
// }

func calcLCM(nums []int64) *big.Int {
	if len(nums) == 0 || nums[0] == 0 {
		return big.NewInt(0)
	}
	lcm := big.NewInt(nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			return big.NewInt(0)
		}

		bigNum := big.NewInt(nums[i])

		absProduct := new(big.Int).Abs(new(big.Int).Mul(lcm, bigNum))
		gcd := calcGCD([]int64{lcm.Int64(), nums[i]})
		lcm = new(big.Int).Div(absProduct, gcd)
	}
	return lcm
}
