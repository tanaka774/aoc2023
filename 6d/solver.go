package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func main() {
	// ans1() // ans:4403592
	ans2() // ans:
}

func ans1() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	ans := 1 // multiply of winning race
	times, distances := parse(scanner)

	records := make([]int, len(times))
	for i, totalT := range times {
		chargeT := int(totalT / 2)

		finding(chargeT, totalT, records, distances, i)
	}

	for _, rec := range records {
		ans *= rec
	}
	fmt.Println("ans:", ans)
}

func ans2() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	totalT, distance := parse2(scanner)

	lowEdge := math.MaxInt
	highEdge := math.MinInt
	// chargeT := int(totalT / 2)
	// previusT := totalT

	// searh from start
	for chargeT := 0; chargeT < totalT; chargeT++ {
		progress := chargeT * (totalT - chargeT)
		if progress > distance {
			lowEdge = chargeT
			break
		}
	}

	// searh from end
	for chargeT := totalT; chargeT > 0; chargeT-- {
		progress := chargeT * (totalT - chargeT)
		if progress > distance {
			highEdge = chargeT
			break
		}
	}
	fmt.Println(lowEdge, highEdge)
	fmt.Println("ans:", highEdge-lowEdge+1)
}

func finding(chargeT int, totalT int, records []int, distances []int, i int) {
	progress := chargeT * (totalT - chargeT)
	if progress <= distances[i] {
		return
	}
	records[i]++
	fmt.Println(i, chargeT)

	if chargeT == totalT/2 {
		finding(chargeT+1, totalT, records, distances, i)
		finding(chargeT-1, totalT, records, distances, i)
	} else if chargeT > totalT/2 {
		finding(chargeT+1, totalT, records, distances, i)
	} else {
		finding(chargeT-1, totalT, records, distances, i)
	}
}

// func finding2(chargeT int, totalT int, preT int, distance int) {
//   if math.Abs(float64(chargeT) - float64(preT)) <= 2
// 	progress := chargeT * (totalT - chargeT)
// 	if progress <= distance {
// 		return
// 	}
// 	(*count)++
//
// 	if chargeT == totalT/2 {
// 		finding2(chargeT+1, totalT, distance, count)
// 		finding2(chargeT-1, totalT, distance, count)
// 	} else if chargeT > totalT/2 {
// 		finding2(chargeT+1, totalT, distance, count)
// 	} else {
// 		finding2(chargeT-1, totalT, distance, count)
// 	}
// }

func parse(scanner *bufio.Scanner) ([]int, []int) {
	times := make([]int, 0)
	distances := make([]int, 0)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		cur := 0
		for {
			if unicode.IsDigit(rune(line[cur])) {
				numStr := getSequenceNumber(line, cur)
				cur += len(numStr)
				num, _ := strconv.Atoi(numStr)
				if i == 0 {
					times = append(times, num)
				} else {
					distances = append(distances, num)
				}
			} else {
				cur++
			}
			if cur >= len(line) {
				break
			}
		}
	}
	return times, distances
}

func parse2(scanner *bufio.Scanner) (int, int) {
	time := ""
	distance := ""
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		cur := 0
		for {
			if unicode.IsDigit(rune(line[cur])) {
				numStr := getSequenceNumber(line, cur)
				cur += len(numStr)
				// num, _ := strconv.Atoi(numStr)
				if i == 0 {
					time += numStr
				} else {
					distance += numStr
				}
			} else {
				cur++
			}
			if cur >= len(line) {
				break
			}
		}
	}
	return atoiEX(time), atoiEX(distance)
}
