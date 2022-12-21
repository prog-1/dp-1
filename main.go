package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func knapsack(s [][]int, n, w, i int, values, weight []int) int {
	if i >= n || w <= 0 {
		return 0
	}
	maxVal1 := 0
	if s[i][w] != 0 {
		return s[i][w]
	}
	if weight[i] <= w {
		maxVal1 = values[i] + knapsack(s, n, w-weight[i], i+1, values, weight)
	}
	maxVal2 := knapsack(s, n, w, i+1, values, weight)
	return max(maxVal1, maxVal2)
}

func knapsack2(n, w int, values, weight []int) int {
	s := make([][]int, n)
	for row := 0; row < n; row++ {
		s[row] = make([]int, w+1)
	}
	return knapsack(s, n, w, 0, values, weight)
}

func main() {
	n := 3
	w := 4
	values := []int{1, 2, 3}
	weight := []int{4, 5, 1}
	fmt.Println(knapsack2(n, w, values, weight))
}
