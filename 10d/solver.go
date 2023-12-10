package main

import (
	"fmt"
)

func main() {
	ans1() // ans:
	// ans2() // ans:
}

func ans1() {
	file, scanner := getScanner("./example.txt")
	// file, scanner = getScanner("./input.txt")
	defer file.Close()

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

	}
}
