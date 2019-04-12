package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Make sure commandline argument is 1 positive integer.")
	}
	result := fibo(n)
	fmt.Println(result)
}

func fibo (n int) int {
	if n == 0 {
		return 0
	}
	fibArr :=  make([]int, n+1)
	fibArr[0] = 0
	fibArr[1] = 1
	for i:=2; i < n + 1; i++ {
		fibArr[i] = fibArr[i-2] + fibArr[i-1]
	}
	return fibArr[n]
}