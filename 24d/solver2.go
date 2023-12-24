package main

import (
	"fmt"
	"regexp"
)

func ans2() {
	file, scanner := getScanner("./example.txt")
	// file, scanner := getScanner("./input.txt")
	defer file.Close()

	hails := make([]Hail, 0)
	for ln := 0; scanner.Scan(); ln++ {
		line := scanner.Text()
		re := regexp.MustCompile(`-?\d+`)
		nums := re.FindAllString(line, -1)
		px, py, pz, vx, vy, vz := atof(nums[0]), atof(nums[1]), atof(nums[2]), atof(nums[3]), atof(nums[4]), atof(nums[5])
		hails = append(hails, Hail{Pos{px, py, pz}, Vel{vx, vy, vz}})
	}
	fmt.Println(hails)

	cnt := 0
	for i := 0; i < len(hails); i++ {
		for j := i + 1; j < len(hails); j++ {
			// y=ax+b
			h1, h2 := hails[i], hails[j]
			a1, a2 := h1.vel.y/h1.vel.x, h2.vel.y/h2.vel.x
			b1, b2 := h1.pos.y-a1*h1.pos.x, h2.pos.y-a2*h2.pos.x
			if a1 == a2 {
				// fmt.Println("parallel", h1, h2)
				continue
			}
			intsx := (b2 - b1) / (a1 - a2)
			intsy := a1*intsx + b1
			// fmt.Println("deb", intsx, intsy, a1, a2, b1, b2, a1-a2, b2-b1, h1, h2)
			if intsx >= testArea[0] && intsx <= testArea[1] && intsy >= testArea[0] && intsy <= testArea[1] {
				if (intsx-h1.pos.x)*h1.vel.x > 0 && (intsx-h2.pos.x)*h2.vel.x > 0 &&
					(intsy-h1.pos.y)*h1.vel.y > 0 && (intsy-h2.pos.y)*h2.vel.y > 0 {
					// fmt.Println("naka", intsx, intsy, h1, h2)
					cnt++
				}
			}
		}
	}
	fmt.Println("ans", cnt)

}
