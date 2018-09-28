package main

import (
	"fmt"
	"strconv"
)

func main() {

	p := []int{1, 2, 3, 5, 7, 9, 11, 13}
	p1, p2 := partition(p)

	printSlice(p1)
	printSlice(p2)
}

func printSlice(p []int) {
	for i := 0; i < len(p); i++ {
		fmt.Print(" " + strconv.Itoa(p[i]))
	}
	fmt.Println()
}
