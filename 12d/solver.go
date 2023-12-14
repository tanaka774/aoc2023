package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	// ans1() // ans:6852
	ans2() // ans:
}

func ans1() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	ans := 0 // sum of arrangements
	for ln := 0; scanner.Scan(); ln++ {
		line := scanner.Text()
		fmt.Println(line)

		spaceIndex := strings.Index(line, " ")
		numStrs := strings.Split(line[spaceIndex+1:], ",")
		damagedNums := make([]int, 0)
		for _, str := range numStrs {

			damagedNums = append(damagedNums, atoiEX(str))
		}
		var excess string
		for i := 0; i < slices.Max(damagedNums)+1; i++ {
			excess += "#"
		}
		springs := line[:spaceIndex]
		ans += dive(springs, damagedNums, "#", excess)
		ans += dive(springs, damagedNums, ".", excess)
		fmt.Println("ans", ans/2)

		// if ln >= 5 {
		// 	break //debug
		// }
	}
}

func ans2() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	ans := 0 // sum of arrangements
	for ln := 0; scanner.Scan(); ln++ {
		line := scanner.Text()
		fmt.Println(line)

		spaceIndex := strings.Index(line, " ")
		numStrs := strings.Split(line[spaceIndex+1:], ",")
		damagedNums := make([]int, 0)
		for _, str := range numStrs {

			damagedNums = append(damagedNums, atoiEX(str))
		}
		aaa := make([]int, 0)
		var bbb string
		for i := 0; i < 5-1; i++ {
			aaa = append(aaa, damagedNums...)
			bbb += line[:spaceIndex] + "?"
		}
		maxContig := slices.Max(damagedNums)
		damagedNums = append(damagedNums, aaa...)
		springs := bbb + line[:spaceIndex]

		fmt.Println(damagedNums)
		fmt.Println(springs)
		var excess string
		for i := 0; i < maxContig+1; i++ {
			excess += "#"
		}
		fmt.Println(excess)
		ans += dive(springs, damagedNums, "#", excess)
		ans += dive(springs, damagedNums, ".", excess)
		fmt.Println("ans", ans/2)

		// if ln >= 1 {
		// 	break //debug
		// }
	}
}

func dive(springs string, damagedNums []int, sp string, excess string) int {
	// if strings.Index(springs, excess) != -1 {
	// 	return 0
	// }
	nums := countSprings(springs)
	// fmt.Println("***", springs)
	// fmt.Println("***", nums)
	if len(nums) > len(damagedNums) {
		return 0
	}
	for i, num := range nums {
		if num != damagedNums[i] {
			// fmt.Println("***", springs)
			// fmt.Println("***", i, nums)
			// fmt.Println("***", i, damagedNums)
			return 0
		}
	}

	count := 0
	if strings.Index(springs, "?") == -1 {
		// fmt.Println("end?", springs)
		// nums := countSprings(springs)
		if len(nums) != len(damagedNums) {
			return 0
		}
		// for i, num := range nums {
		// 	if num != damagedNums[i] {
		// 		return 0
		// 	}
		// }
		// fmt.Println("valid?", springs)
		// fmt.Println("valid?", nums)
		return 1 //TODO dont return twice on a condition
	}
	springs = strings.Replace(springs, "?", sp, 1)

	count += dive(springs, damagedNums, "#", excess)
	count += dive(springs, damagedNums, ".", excess)
	return count
}

func countSprings(str string) []int {
	// count until \s
	damagedNums := make([]int, 0)
	damagedNum := 0
	for i, ch := range str {
		// dogshit logic
		switch ch {
		case '?':
			// fmt.Println("?maji?")
			return damagedNums
		case '#':
			damagedNum++
		case '.':
			if damagedNum >= 1 {
				damagedNums = append(damagedNums, damagedNum)
				damagedNum = 0
			}
			// case ' ': // unneessary maybe
			// 	if damagedNum >= 1 {
			// 		damagedNums = append(damagedNums, damagedNum)
			// 		damagedNum = 0
			// 	}
			// 	break
		}
		if i == len(str)-1 {
			if damagedNum >= 1 {
				damagedNums = append(damagedNums, damagedNum)
				damagedNum = 0
			}
			break
		}
	}

	return damagedNums
}
