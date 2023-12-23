package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

const (
	MAX_X = 3 //ex
	MAX_Y = 3
	// MAX_X = 10 //input
	// MAX_Y = 10
	EMPTY = -1
	CUBE  = 1
)

type Space = map[int][][]int

func main() {
	// f, err := os.Open("example.txt")
	f, err := os.Open("input.txt") // ans1:
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// scanner.Split(bufio.ScanWords) // get each char

	// space := make(Space, 0)
	// field := make([][]int, 0)
	cubes := make([][][]int, 0) // [[[x1,y1,z1], [x2,y2,z2]], ...]
	re := regexp.MustCompile(`\d+`)
	for ln := 0; scanner.Scan(); ln++ {
		// fmt.Println(scanner.Text())
		line := scanner.Text()
		nums := re.FindAllString(line, -1)
		x1, y1, z1, x2, y2, z2 := atoi(nums[0]), atoi(nums[1]), atoi(nums[2]), atoi(nums[3]), atoi(nums[4]), atoi(nums[5])
		// if x1 != x2 {
		// 	field := getField(space, MAX_Y, MAX_X, z1)
		// 	for x := min(x1, x2); x <= max(x1, x2); x++ {
		// 		field[y1][x] = CUBE + ln
		// 	}
		// 	space[z1] = slices.Clone(field)
		// }
		// if y1 != y2 {
		// 	field := getField(space, MAX_Y, MAX_X, z1)
		// 	for y := min(y1, y2); y <= max(y1, y2); y++ {
		// 		field[y][x1] = CUBE + ln
		// 	}
		// 	space[z1] = slices.Clone(field)
		// }
		// if z1 != z2 {
		// 	for z := min(z1, z2); z <= max(z1, z2); z++ {
		// 		field := getField(space, MAX_Y, MAX_X, z)
		// 		field[y1][x1] = CUBE + ln
		// 		space[z] = slices.Clone(field)
		// 	}
		// }

		cubes = append(cubes, [][]int{{x1, y1, z1}, {x2, y2, z2}})
	}

	// fmt.Println(space)
	fmt.Println(cubes)

	slices.SortFunc(cubes, func(a, b [][]int) int {
		if a[0][2] < b[0][2] && a[1][2] > b[0][2] {
			fmt.Println("???", a, b)
		}
		return a[0][2] - b[0][2]
	})
	fmt.Println(cubes)

	for i1 := len(cubes)-1; i1>=0; i1-- {
		cube := cubes[i1]
		for x := cube[0][0]; x < cube[1][0]; x++ {
			for y := cube[0][1]; y < cube[1][1]; y++ {
				for z := cube[0][2]; z < cube[1][2]; z++ {
					if 

					// for i2 := range cubes {
					// 	cube2 := cubes[len(cubes)-1-i2]

					// }
				}
			}
		}
	}
}

func getField(space Space, maxY, maxX, z int) [][]int {
	for k := range space {
		if z == k {
			return space[k]
		}
	}

	a := make([][]int, maxY)
	for i := range a {
		b := make([]int, maxX)
		for i2 := range b {
			b[i2] = EMPTY
		}
		a[i] = slices.Clone(b)
	}
	return a
}

func atoi(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}
