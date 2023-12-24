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
	// MAX_X = 3 //ex
	// MAX_Y = 3
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
	cubesMap := make(map[int][][][]int)
	// field := make([][]int, 0)
	// cubes := make([][][]int, 0) // [[[x1,y1,z1], [x2,y2,z2]], ...]
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
		// 		space[z] = slices.Clone(field)space
		// 	}
		// }

		// cubes = append(cubes, [][]int{{x1, y1, z1}, {x2, y2, z2}})
		// if you need, label each cube with line number for telling the difference
		// actually, there is no need to store z value here, maybe?
		if z1 == z2 {
			cubesMap[z1] = append(cubesMap[z1], [][]int{{x1, y1, z1, ln}, {x2, y2, z1, ln}})
		} else {
			for z := z1; z <= z2; z++ {
				cubesMap[z] = append(cubesMap[z], [][]int{{x1, y1, z, ln}, {x2, y2, z, ln}})
			}
		}
	}

	// fmt.Println(space)
	// fmt.Println(cubes)
	//
	// slices.SortFunc(cubes, func(a, b [][]int) int {
	// 	if a[0][2] < b[0][2] && a[1][2] > b[0][2] {
	// 		fmt.Println("???", a, b)
	// 	}
	// 	return a[0][2] - b[0][2]
	// })
	// fmt.Println(cubes)

	// fmt.Println(cubesMap)
	for k, v := range cubesMap {
		fmt.Println(k, v[0][0][2])
	}

	// first, let them fall down
	for {
		isFalledOnce := false

		for k := range copyMap(cubesMap) {
			for k2, v := range cubesMap { // debug
				fmt.Println(k2, v[0][0])
				if k2 != v[0][0][2] {
					fmt.Println("???", k2, v)
				} else if v[0][0][2] < 0 {
					fmt.Println("why", k2, v)
				}
			}

			if k == 1 {
				continue
			}
			isFalled := false

			if len(cubesMap[k-1]) == 0 {
				for i := range slices.Clone(cubesMap[k]) {
					fmt.Println("len0", k, cubesMap[k][i], cubesMap[k][i][0][2])
					cubesMap[k][i][0][2]--
					cubesMap[k][i][1][2]--
					isFalled = true
				}
			}

			for i, cubes := range slices.Clone(cubesMap[k]) {
				for _, downCubes := range slices.Clone(cubesMap[k-1]) {
					if cubes[0][0] <= downCubes[1][0] && cubes[1][0] >= downCubes[0][0] &&
						cubes[0][1] <= downCubes[1][1] && cubes[1][1] >= downCubes[0][1] {
						continue // block below
					}
					fmt.Println("intheair", k, cubesMap[k][i], cubesMap[k][i][0][2])
					cubesMap[k][i][0][2]--
					cubesMap[k][i][1][2]--
					isFalled = true
				}
			}

			if isFalled {
				temp := cubesMap[k]
				delete(cubesMap, k)
				cubesMap[k-1] = append(cubesMap[k-1], temp...)
				isFalledOnce = true
			}
		}

		if !isFalledOnce {
			break
		}
	}
	fmt.Println(cubesMap)
	for k, v := range cubesMap {
		fmt.Println(k, v[0][0][2])
	}

	cnt := 0 // the number of "disintegrate"
	for k := range cubesMap {
		if len(cubesMap[k+1]) == 0 {
			fmt.Println("top", len(cubesMap[k]))
			cnt += len(cubesMap[k])
			continue
		}
		for _, cubes := range cubesMap[k] {
			canDisintegrate := true
			hasNonAbove := true
			zCubes := false
			// every cube above needs to have supports for disintegrating
			for _, aboveCubes := range cubesMap[k+1] {
				if cubes[0][2] != aboveCubes[0][2] && cubes[0][3] == aboveCubes[0][3] {
					zCubes = true
					fmt.Println("zcubes", "above", aboveCubes, "down", cubes)
					continue
				}
				if !(cubes[0][0] <= aboveCubes[1][0] && cubes[1][0] >= aboveCubes[0][0] &&
					cubes[0][1] <= aboveCubes[1][1] && cubes[1][1] >= aboveCubes[0][1] &&
					cubes[0][3] != aboveCubes[0][3]) {
					fmt.Println("out", "above", aboveCubes, "down", cubes)
					continue
				}
				fmt.Println("above", aboveCubes, "down", cubes)
				// if cubes[0][3] != aboveCubes[0][3] {
				hasNonAbove = false
				// }
				hasOtherSupport := false
				// check whether cube above has other supporting cube
				for _, otherCubes := range cubesMap[k] {
					if cubes[0][3] == otherCubes[0][3] {
						continue // same cubes
					}

					if otherCubes[0][0] <= aboveCubes[1][0] && otherCubes[1][0] >= aboveCubes[0][0] &&
						otherCubes[0][1] <= aboveCubes[1][1] && otherCubes[1][1] >= aboveCubes[0][1] &&
						otherCubes[0][3] != aboveCubes[0][3] {
						hasOtherSupport = true
						break
					}
				}
				if !hasOtherSupport {
					canDisintegrate = false
				}
			}
			if (hasNonAbove || canDisintegrate) && !zCubes {
				fmt.Println("disint", cubes, hasNonAbove, canDisintegrate)
				cnt++
			}
		}
	}
	fmt.Println(cnt)
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

func copyMap[T comparable, R any](m map[T]R) map[T]R {
	temp := make(map[T]R)
	for k, v := range m {
		temp[k] = v
	}
	return temp
}
