package main

import "fmt"

func main() {
	/*TODO: Hard-coded input needs to go*/
	n := 6
	m := 7
	//edgeList contains edges only once, we assume the given graph is an undirected graph
	edgeList := [][]int{
		{1, 2},
		{2, 3},
		{6, 3},
		{5, 6},
		{2, 5},
		{2, 4},
		{4, 1}}
	fmt.Println(deg(n, m, edgeList))
}

func deg(n int, m int, edgeList [][]int) []int {
	results := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < 2; j++ {
			results[edgeList[i][j]-1] = results[edgeList[i][j]-1] + 1
		}
	}
	return results
}
