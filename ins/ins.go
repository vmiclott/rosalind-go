package main

import "fmt"

func main() {
	/*TODO: Hard-coded input needs to go*/
	n := 6
	a := []int{6, 10, 4, 5, 1, 2}
	fmt.Println(ins(n, a))
	fmt.Println(a)
}

//Insertion-sort a (length n)
func ins(n int, a []int) int {
	count := 0
	for i := 1; i < n; i++ {
		k := i
		for k > 0 && a[k] < a[k-1] {
			temp := a[k-1]
			a[k-1] = a[k]
			a[k] = temp
			k = k - 1
			count = count + 1
		}
	}
	return count
}
