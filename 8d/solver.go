package main

import (
	"fmt"
)

func main() {
	// ans1() // ans:19951
	ans2() // ans:16342438708751
}

func ans1() {
	file, scanner := getScanner("./example.txt")
	// file, scanner := getScanner("./input.txt")
	defer file.Close()

	opes := ""
	nodeHelper := make(map[byte][]string, 0)
	nodes := make(map[string][]string, 0)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		if i == 0 {
			opes = line
			continue
		}
		if i == 1 {
			continue
		}

		nodeHelper[line[0]] = append(nodeHelper[line[0]], line[:3])
		nodes[line[:3]] = []string{line[7:10], line[12:15]}
	}

	cur := "AAA"
	step := 0
	var next string //unnessesary?
	for {
		for _, ope := range opes {
			if ope == 'L' {
				next = nodes[cur][0]
			} else {
				next = nodes[cur][1]
			}
			step++
			if next == "ZZZ" {
				break
			}
			cur = next
		}
		if next == "ZZZ" {
			break
		}
	}
	fmt.Println("ans:", step)
}

func ans2() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	opes := ""
	nodeHelper := make(map[byte][]string, 0)
	nodes := make(map[string][]string, 0)
	curs := make([]string, 0)
	goals := make([]string, 0)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		if i == 0 {
			opes = line
			continue
		}
		if i == 1 {
			continue
		}

		nodeHelper[line[0]] = append(nodeHelper[line[0]], line[:3])
		nodes[line[:3]] = []string{line[7:10], line[12:15]}

		if line[2] == 'A' {
			curs = append(curs, line[:3])
		} else if line[2] == 'Z' {
			goals = append(goals, line[:3])
		}
	}

	history := make([][]string, len(curs))
	loopCount_deb := make([][]int, len(curs))
	loopEnd := false
Loop:
	for {
		for _, ope := range opes {
			for curIndex, cur := range curs {
				history[curIndex] = append(history[curIndex], cur)
				if cur[2] == 'Z' {
					if len(loopCount_deb[curIndex]) < 3 {
						loopCount_deb[curIndex] = append(loopCount_deb[curIndex], len(history[curIndex]))
					}
					history[curIndex] = make([]string, 0)
				}
				if ope == 'L' {
					curs[curIndex] = nodes[cur][0]
				} else {
					curs[curIndex] = nodes[cur][1]
				}
			}
			loopEnd = every[[]int](loopCount_deb, func(i []int) bool {
				return len(i) == 3
			})
			if loopEnd {
				break Loop
			}
		}
		// if loopEnd {
		// 	break
		// }
	}
	loopCount := make([]int64, len(curs))
	for i := range loopCount_deb {
		loopCount[i] = int64(loopCount_deb[i][2])
	}
	lcm := calcLCM(loopCount)
	fmt.Println("ans:", lcm)
}
