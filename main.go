package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func knapsack(v, w []int, mw, n int, cache map[int]int) int {
	if v, ok := cache[n]; ok {
		return v
	}
	if n == 0 || mw == 0 {
		return 0
	}
	maxValue := 0
	if w[n-1] > mw {
		maxValue = knapsack(v, w, mw, n-1, cache)
	} else {
		maxValue = max(v[n-1]+knapsack(v, w, mw-w[n-1], n-1, cache), knapsack(v, w, mw, n-1, cache))
	}
	cache[n] = maxValue
	return maxValue
}
func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Enter N and W:")
	var n, w int
	fmt.Fscan(r, &n, &w)
	values := make([]int, n)
	weight := make([]int, n)
	fmt.Println("Enter values:")
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &values[i])
	}
	fmt.Println("Enter weight:")
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &weight[i])
	}

	fmt.Println(knapsack(values, weight, w, n, map[int]int{}))
}
