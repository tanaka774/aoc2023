package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	// ans1() // ans:253933213
	ans2() // ans:253473930
}

func ans1() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	ans := 0 // sum of rakn*bid
	points := make([]int, 0)
	pointBid := make(map[int]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		hands := line[:5]
		bid := atoiEX(line[6:])

		point := calcTypePoint(hands) + calcLabelPoint(hands)

		points = append(points, point)
		pointBid[point] = bid
	}

	slices.Sort(points)
	for i, point := range points {
		ans += pointBid[point] * (i + 1)
	}
	fmt.Println("ans:", ans)
}

func ans2() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	ans := 0 // sum of rakn*bid
	points := make([]int, 0)
	pointBid := make(map[int]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		hands := line[:5]
		bid := atoiEX(line[6:])

		point := calcTypePoint2(hands) + calcLabelPoint2(hands)

		points = append(points, point)
		pointBid[point] = bid
	}

	slices.Sort(points)
	for i, point := range points {
		ans += pointBid[point] * (i + 1)
	}
	fmt.Println("ans:", ans)
}

func calcTypePoint(str string) int {
	sorted := sortString(str)
	same := 1
	sames := make([]int, 0)
	for i := range sorted {
		if i >= len(sorted)-1 {
			if same >= 2 {
				sames = append(sames, same)
			}
			break
		}
		if sorted[i] == sorted[i+1] {
			same++
		} else {
			if same >= 2 {
				sames = append(sames, same)
				same = 1
			}
		}
	}

	if len(sames) == 0 {
		return 200000000
	} else if sames[0] == 5 {
		return 800000000
	} else if sames[0] == 4 {
		return 700000000
	} else if len(sames) >= 2 && sames[0]+sames[1] == 5 {
		// fullhouse
		return 600000000
	} else if sames[0] == 3 {
		return 500000000
	} else if len(sames) >= 2 && sames[0]+sames[1] == 4 {
		// two pair
		return 400000000
	} else {
		// one pair
		return 300000000
	}
}

func calcLabelPoint(str string) int {
	point := 0
	cards := []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	for i, ch := range str {
		for j, card := range cards {
			if ch == card {
				point += int(math.Pow(float64(len(cards)), float64(len(str)-1-i))) * (len(cards) - j)
			}
		}
	}
	return point
}

func calcTypePoint2(str string) int {
	sorted := sortString(str)
	same := 1
	sames := make([]int, 0)
	jCount := 0
	for i := range sorted {
		if 'J' == sorted[i] {
			jCount++
			continue
		}

		if i >= len(sorted)-1 {
			if same >= 2 {
				sames = append(sames, same)
			}
			break
		}
		if sorted[i] == sorted[i+1] {
			same++
		} else {
			if same >= 2 {
				sames = append(sames, same)
				same = 1
			}
		}
	}

	if jCount == 5 {
		sames = append(sames, 5)
	} else if len(sames) == 0 && jCount >= 1 {
		sames = append(sames, 1+jCount)
	} else if jCount >= 1 {
		maxIndex := slices.Index(sames, slices.Max(sames))
		sames[maxIndex] += jCount
	}

	if len(sames) == 0 {
		return 20000000
	} else if sames[0] == 5 {
		return 80000000
	} else if sames[0] == 4 {
		return 70000000
	} else if len(sames) >= 2 && sames[0]+sames[1] == 5 {
		// fullhouse
		return 60000000
	} else if sames[0] == 3 {
		return 50000000
	} else if len(sames) >= 2 && sames[0]+sames[1] == 4 {
		// two pair
		return 40000000
	} else {
		// one pair
		return 30000000
	}
}

func calcLabelPoint2(str string) int {
	point := 0
	cards := []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
	for i, ch := range str {
		for j, card := range cards {
			if ch == card {
				point += int(math.Pow(float64(len(cards)), float64(len(str)-1-i))) * (len(cards) - j)
			}
		}
	}
	return point
}
