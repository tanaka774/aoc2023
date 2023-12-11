package main

import (
	"fmt"
	"slices"
	"strings"
)

const (
	BOOM   = 9999
	DONT   = 999
	UNINIT = -9999
	NORTH  = 0
	SOUTH  = 1
	EAST   = 2
	WEST   = 3
)

type stringSet map[string]bool

func main() {
	// ans1() // ans:6870
	ans2() // ans:
}

func ans1() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	field := make([]string, 0)
	sp := make([]int, 2) // y,x
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		field = append(field, line)
		if strings.Index(line, "S") != -1 {
			sp[0] = i
			sp[1] = strings.Index(line, "S")
		}
	}

	cp1 := sp
	cp2 := sp
	step := 1     // if you set position one step forward
	var dire1 int // north:0, south:1, east:2, west:3
	var dire2 int // north:0, south:1, east:2, west:3

	// for actual input
	cp1 = []int{sp[0] + 1, sp[1]}
	dire1 = NORTH
	cp2 = []int{sp[0] - 1, sp[1]}
	dire2 = SOUTH

	// // for example
	// cp1 = []int{sp[0] + 1, sp[1]}
	// dire1 = NORTH
	// cp2 = []int{sp[0] - 1, sp[1]}
	// dire2 = SOUTH

	for {
		step++
		cp1[0], cp1[1], dire1 = goon(field, dire1, cp1[0], cp1[1])
		cp2[0], cp2[1], dire2 = goon(field, dire2, cp2[0], cp2[1])
		if cp1[0] == cp2[0] && cp1[1] == cp2[1] {
			fmt.Println("end(ans):", step)
			break
		}
	}
}

func ans2() {
	// file, scanner := getScanner("./example2-1.txt")
	// file, scanner := getScanner("./example2-2.txt")
	// file, scanner := getScanner("./example2-3.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	field := make([]string, 0)
	sp := make([]int, 2) // y,x
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		field = append(field, line)
		if strings.Index(line, "S") != -1 {
			sp[0] = i
			sp[1] = strings.Index(line, "S")

		}
	}

	cp := sp     // ok not returning copy?
	var dire int // north:0, south:1, east:2, west:3
	loops := make([][][]int, 0)
	loops_set := make(stringSet, 0)
	unconnected := make(stringSet, 0)

	// for example2-1,2,3
	cp = []int{sp[0] + 1, sp[1]}
	dire = NORTH

	loopArr := make([][]int, 0)
	for {
		cp[0], cp[1], dire = goon2(field, dire, cp[0], cp[1])
		loopArr = append(loopArr, slices.Clone(cp))
		if cp[0] == sp[0] && cp[1] == sp[1] {
			loops = append(loops, loopArr)

			for _, lo := range loopArr {
				str := fmt.Sprintf("%d,%d", lo[0], lo[1])
				loops_set[str] = true
			}

			loopArr = make([][]int, 0)
			break
		}
	}

	for fy, row := range field {
		for fx, ch := range row {
			str := fmt.Sprintf("%d,%d", fy, fx)
			if ch == '.' {
				unconnected[str] = true
			}
			if loops_set[str] || unconnected[str] {
				continue
			}
			dire := decideDire(ch)
			searched, isLoop := goaround(field, dire, fy, fx, loops_set, unconnected)
			for _, se := range searched {
				str := fmt.Sprintf("%d,%d", se[0], se[1])
				if isLoop {
					loops_set[str] = true
					loops = append(loops, searched)
				} else {
					unconnected[str] = true
				}
			}
		}
	}
	fmt.Println(len(loops))
	fmt.Println(len(loops_set))
	// fmt.Println(unconnected)

	allun := make(stringSet)
	for k, v := range unconnected {
		allun[k] = v
	}
	for unc := range unconnected {
		yx := strings.Split(unc, ",")
		y := atoiEX(yx[0])
		x := atoiEX(yx[1])

		went := make(stringSet, 0)
		isOpen := false
		getOutOfPipes(field, UNINIT, y, x, went, &isOpen, unconnected)
		fmt.Println("went", went)
		// fmt.Println("allun", allun)
		if isOpen {
			for k := range went {
				delete(allun, k)
			}
			// fmt.Println("allun(isopen)", allun)
		}
	}

	fmt.Println("unc", unconnected)
	fmt.Println("allun", allun)
	printArray[[][]int](loops)
	fmt.Println("ans?:", len(allun))
}

func getOutOfPipes(field []string, dire, cy, cx int, went stringSet, isOpen *bool, unconnected stringSet) {
	if *isOpen || cy < 0 || cy > len(field)-1 || cx < 0 || cx > len(field[0])-1 {
		// here is open, delete keys of went from unconnected
		*isOpen = true
		return //?
	}

	ndire := getDire2(field, cy, cx, dire, unconnected)
	if ndire == BOOM {
		return
	}

	str := fmt.Sprintf("%d,%d", cy, cx)
	if went[str] {
		return //?
	}
	went[str] = true

	getOutOfPipes(field, NORTH, cy+1, cx, went, isOpen, unconnected)
	getOutOfPipes(field, SOUTH, cy-1, cx, went, isOpen, unconnected)
	getOutOfPipes(field, WEST, cy, cx+1, went, isOpen, unconnected)
	getOutOfPipes(field, EAST, cy, cx-1, went, isOpen, unconnected)
}

func goaround(field []string, dire, cy, cx int, loops_set, unconncted stringSet) (searched [][]int, isLoop bool) {
	sy, sx := cy, cx
	isLoop = false
	for {
		searched = append(searched, []int{cy, cx})
		cy, cx, dire = goon2(field, dire, cy, cx)
		str := fmt.Sprintf("%d,%d", cy, cx)
		if loops_set[str] || unconncted[str] ||
			cy < 0 || cy > len(field)-1 || cx < 0 || cx > len(field[0])-1 {
			break
		}

		if sy == cy && sx == cx {
			if len(searched) >= 2 {
				isLoop = true
				break
			} else {
				break
			}
		}
	}
	return
}

/**
* return one of the direction according to tile
 */
func decideDire(tile rune) int {
	switch tile {
	case '|':
		return SOUTH
	case '-':
		return WEST
	case 'L':
		return WEST
		// return NORTH
	case 'J':
		return EAST
		// return NORTH
	case '7':
		return NORTH
		// return SOUTH
	case 'F':
		return NORTH
		// return SOUTH
	}
	return BOOM // '.' should come
}

func goon(field []string, dire int, cy int, cx int) (ny, nx, ndire int) {
	tileMap := map[byte][][]int{
		'|': {{1, 0}, {-1, 0}, {DONT}, {DONT}},  // from n, s
		'-': {{DONT}, {DONT}, {0, -1}, {0, 1}},  // e, w
		'L': {{0, 1}, {DONT}, {-1, 0}, {DONT}},  // n, e
		'J': {{0, -1}, {DONT}, {DONT}, {-1, 0}}, // n, w
		'7': {{DONT}, {0, -1}, {DONT}, {1, 0}},  // s, w
		'F': {{DONT}, {0, 1}, {1, 0}, {DONT}},   // s, e
	}
	tile := field[cy][cx]
	if tileMap[tile][dire][0] == DONT {
		ny, nx, ndire = cy, cx, dire
		return
	}
	ny = cy + tileMap[tile][dire][0]
	nx = cx + tileMap[tile][dire][1]
	ndire = getDire(tile, dire)
	return
}

func goon2(field []string, dire int, cy int, cx int) (ny, nx, ndire int) {
	tileMap := map[byte][][]int{
		'|': {{1, 0}, {-1, 0}, {DONT}, {DONT}},  // from n, s
		'-': {{DONT}, {DONT}, {0, -1}, {0, 1}},  // e, w
		'L': {{0, 1}, {DONT}, {-1, 0}, {DONT}},  // n, e
		'J': {{0, -1}, {DONT}, {DONT}, {-1, 0}}, // n, w
		'7': {{DONT}, {0, -1}, {DONT}, {1, 0}},  // s, w
		'F': {{DONT}, {0, 1}, {1, 0}, {DONT}},   // s, e
	}
	tile := field[cy][cx]
	if tileMap[tile][dire][0] == DONT {
		ny, nx, ndire = cy, cx, dire
		return
	}
	ny = cy + tileMap[tile][dire][0]
	nx = cx + tileMap[tile][dire][1]
	ndire = getDire(tile, dire)
	return
}

func getDire(tile byte, dire int) int {
	direMap := map[byte][]int{
		'|': {NORTH, SOUTH, BOOM, BOOM}, // from n, s
		'-': {BOOM, BOOM, EAST, WEST},   // e, w
		'L': {WEST, BOOM, SOUTH, BOOM},  // n, e
		'J': {EAST, BOOM, BOOM, SOUTH},  // n, w
		'7': {BOOM, EAST, BOOM, NORTH},  // s, w
		'F': {BOOM, WEST, NORTH, BOOM},  // s, e
	}
	return direMap[tile][dire]
}

func getDire2(field []string, cy, cx int, dire int, unconnected stringSet) int {
	for unc := range unconnected { // unnecessary???
		yx := strings.Split(unc, ",")
		y := atoiEX(yx[0])
		x := atoiEX(yx[1])
		if cy == y && cx == x {
			return UNINIT
		}
	}
	tile := field[cy][cx]
	if dire == UNINIT || tile == '.' {
		return UNINIT
	}
	if tile == 'S' {
		return BOOM
	}
	direMap := map[byte][]int{
		'|': {NORTH, SOUTH, BOOM, BOOM}, // from n, s
		'-': {BOOM, BOOM, EAST, WEST},   // e, w
		'L': {WEST, BOOM, SOUTH, BOOM},  // n, e
		'J': {EAST, BOOM, BOOM, SOUTH},  // n, w
		'7': {BOOM, EAST, BOOM, NORTH},  // s, w
		'F': {BOOM, WEST, NORTH, BOOM},  // s, e
	}
	return direMap[tile][dire]
}
