package main

import (
	"fmt"
)

func main() {
	/*TODO: Hard-coded input needs to go*/
	a := []int{10, 20, 30, 40, 50}
	k := []int{40, 10, 35, 15, 40, 20}
	m := len(k)
	results := make([]int, m)
	for i := 0; i < m; i++ {
		results[i] = bins(a, k[i], 0)
	}
	fmt.Println("RESULT: ", results)
}

func bins(a []int, s int, startIndex int) int {
	n := len(a)
	if n == 1 && a[0] == s {
		return startIndex + 1 // + 1 because no 0-indexing
	} else if n == 1 {
		return -1
	}
	if s < a[n/2] {
		return bins(a[:n/2], s, startIndex)
	} else {
		return bins(a[n/2:], s, startIndex+n/2)
	}
}
