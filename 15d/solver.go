package main

import (
	"fmt"
)

func main() {
	ans1() // ans:
	// ans2() // ans:
}

func ans1() {
	// file, scanner := getScanner("./example.txt")
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	ans := 0
	for ln := 0; scanner.Scan(); ln++ {
		line := scanner.Text()
		str := ""
		for i, ch := range line {
			if i == len(line)-1 {
				str += string(ch)
				ans += hash(str)
				// fmt.Println(str, hash(str))
			} else if ch == ',' {
				ans += hash(str)
				// fmt.Println(str, hash(str))
				str = ""
			} else {
				str += string(ch)
			}
		}
	}
	fmt.Println(ans)
}

func hash(str string) (res int) {
	for _, ch := range str {
		res += int(ch)
		res *= 17
		res %= 256
	}
	return
}
