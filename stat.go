package main

import "sort"

func partition(p []int) (p1 []int, p2 []int) {

	sort.Ints(p)
	count := len(p)

	medianIdx := count / 2
	p1 = p[0:medianIdx]
	p2 = p[medianIdx:count]
	return
}
