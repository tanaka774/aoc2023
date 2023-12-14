package main

import (
	"fmt"
	"strings"
)

const (
// EOF_LN = 15 // 1-ex
// EOF_LN = 1339 // 1-input
)

func main() {
	ans1() // ans:
	// ans2() // ans:
}

func ans1() {
	// file, scanner := getScanner("./example.txt")
	// EOF_LN := 15 // 1-ex
	file, scanner := getScanner("./input.txt")
	EOF_LN := 1339 // 1-input
	defer file.Close()

	mountain := make([]string, 0)
	ans := 0 // sum of divider

	for ln := 0; scanner.Scan(); ln++ {
		line := scanner.Text()

		mountain = append(mountain, line)

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err, "ln", ln)
		}

		if line == "" || ln == EOF_LN-1 {
			printArray(mountain, "mo")
			row := findDividerUp(mountain)
			fmt.Println("up row", row)
			rowd := findDividerDown(mountain)
			fmt.Println("d row", rowd)
			re := reverse(mountain)
			printArray(re, "rev")
			col := findDividerUp(re)
			fmt.Println("up col", col)
			ans += row*100 + col
			cold := findDividerDown(re)
			fmt.Println("d col", cold)
			ans += rowd*100 + cold

			mountain = make([]string, 0)

			// if ln >= 140 {
			// 	break // debug
			// }
		}

	}
	fmt.Println("ans", ans)
}

func findDividerUp(mountain []string) int {
	mid := len(mountain) / 2
	// Loop:
	for r1 := 1; r1 <= mid; r1++ {
		// fmt.Println(mountain[r1])
		for r2 := 0; r2 < mid*2; r2++ {
			if mid+r1+r2 > len(mountain)-1 || mid+r1-1-r2 < 0 {
				fmt.Println("up", mid, r1, r2)
				return mid + r1
			}
			fmt.Println("***", mid+r1+r2, mountain[mid+r1+r2])
			fmt.Println("***", mid+r1-1-r2, mountain[mid+r1-1-r2])
			fmt.Println("")
			if !strings.EqualFold(mountain[mid+r1+r2], mountain[mid+r1-1-r2]) {
				// fmt.Println("***", mid+r1+r2, mountain[mid+r1+r2])
				// fmt.Println("***", mid+r1-1-r2, mountain[mid+r1-1-r2])
				// fmt.Println("")
				if mid+r1+r2 == 6 && mid+r1-1-r2 == 5 {
					fmt.Println(mountain[mid+r1+r2])
					fmt.Println(mountain[mid+r1-1-r2])
				}
				break
			}
		}

		// for r2 := 0; r2 < mid*2; r2++ {
		// 	if mid-r1+1+r2 > len(mountain)-1 || mid-r1-r2 < 0 {
		// 		fmt.Println(mid, r1, r2)
		// 		return mid - r1 + 1
		// 	}
		// 	if !strings.EqualFold(mountain[mid-r1+1+r2], mountain[mid-r1-r2]) {
		// 		break
		// 	}
		// }
	}
	return 0
}

func findDividerDown(mountain []string) int {
	mid := len(mountain) / 2
	// Loop:
	for r1 := 1; r1 <= mid; r1++ {
		for r2 := 0; r2 < mid*2; r2++ {
			if mid-r1+1+r2 > len(mountain)-1 || mid-r1-r2 < 0 {
				fmt.Println("down", mid, r1, r2)
				return mid - r1 + 1
			}
			fmt.Println("***", mid-r1+1+r2, mountain[mid-r1+1+r2])
			fmt.Println("***", mid-r1-r2, mountain[mid-r1-r2])
			fmt.Println("")
			if !strings.EqualFold(mountain[mid-r1+1+r2], mountain[mid-r1-r2]) {
				break
			}
		}
	}
	return 0
}

func reverse(mountain []string) []string {
	reborn := make([]string, len(mountain[0]))
	for _, row := range mountain {
		for j, ch := range row {
			reborn[j] += string(ch)
		}
	}
	return reborn
}
