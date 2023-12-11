package main

import (
	"fmt"
)

func printArray[T any](arr []T) {
	// TODO better visual
	fmt.Println("--- [")
	for _, ele := range arr {
		fmt.Println(ele)
	}
	fmt.Println("] ---")
}
