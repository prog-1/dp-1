package main

import (
	"fmt"
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
	n := 3
	w := 1000
	values := []int{1, 2, 3, 8}
	weight := []int{4, 5, 1, 1}
	fmt.Println(knapsack(values, weight, w, n, map[int]int{}))
}
