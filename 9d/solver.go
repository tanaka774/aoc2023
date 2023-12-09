package main

import (
	"fmt"
)

func main() {
	// ans1() // ans:2043677056
	ans2() // ans:1062
}

func ans1() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	history := make([][]int, 0)
	ans := 0 // sum of each last number

	for scanner.Scan() {
		line := scanner.Text()
		nums := getNumsWithMinus(line, 0)

		history = append(history, nums)

		for {
			nexts := make([]int, 0)
			temp := history[len(history)-1]
			for i := range temp {
				if i == 0 {
					continue
				}
				nexts = append(nexts, temp[i]-temp[i-1])
			}
			history = append(history, nexts)
			analyzeEnd := every[int](nexts, func(i int) bool {
				return i == 0
			})
			if analyzeEnd {
				break
			}
		}
		fmt.Println(history)
	}
	for _, his := range history {
		ans += his[len(his)-1]
	}

	fmt.Println("ans:", ans)
}

func ans2() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	history := make([][]int, 0)
	ans := 0 // sum of each last number

	for scanner.Scan() {
		line := scanner.Text()
		nums := getNumsWithMinus(line, 0)

		history = append(history, nums)

		for {
			nexts := make([]int, 0)
			temp := history[len(history)-1]
			for i := range temp {
				if i == 0 {
					continue
				}
				nexts = append(nexts, temp[i]-temp[i-1])
			}
			history = append(history, nexts)
			analyzeEnd := every[int](nexts, func(i int) bool {
				return i == 0
			})
			if analyzeEnd {
				break
			}
		}
		// fmt.Println(history)
	}

	rec := 0
	for i, his := range history {
		if i > 0 && len(his) > len(history[i-1]) {
			rec = 0
		}
		if rec%2 == 0 {
			ans += his[0]
		} else {
			ans -= his[0]
		}
		rec++
	}

	fmt.Println("ans:", ans)
}

// try to use history as [][][]int properly
func ans2_DISCARDED() {
	file, scanner := getScanner("./example.txt")
	// file, scanner := getScanner("./input.txt")
	defer file.Close()

	history := make([][][]int, 0)
	ans := 0 // sum of each last number

	for ln := 0; scanner.Scan(); ln++ {
		line := scanner.Text()
		nums := getNumsWithMinus(line, 0)

		tempArr := make([][]int, 0)
		tempArr = append(tempArr, nums)
		history = append(history, tempArr)

		for {
			nexts := make([][]int, 0)
			temp := history[ln][len(history)-1]
			for i := range temp {
				if i == 0 {
					continue
				}
				tempArr := make([]int, 0)
				tempArr = append(tempArr, temp[i]-temp[i-1])
				nexts = append(nexts, tempArr)
			}
			history[ln] = append(history[ln], nexts[ln])
			fmt.Println(history)
			analyzeEnd := every[int](nexts[ln], func(i int) bool {
				return i == 0
			})
			if analyzeEnd {
				break
			}
		}
		fmt.Println(history)
	}

	// for i, his := range history {
	// 	if i == 0 {
	// 		ans += his[0]
	// 	} else {
	// 		ans -= his[0]
	// 	}
	// 	fmt.Println("ans:", ans)
	// }

	fmt.Println("ans:", ans)
}
