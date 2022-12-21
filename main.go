package main

import "fmt"

func knapsack(n, w int, val, wt []int) int {
	table := make([][]int, n)
	for i := range table {
		table[i] = make([]int, w+1)
	}
	return recursion(table, n, w, 0, val, wt)
}

func recursion(table [][]int, n, w, curr int, val, wt []int) int {
	var maxVal1, maxVal2 int
	if w <= 0 || n <= 0 || curr >= n || curr < 0 {
		return 0
	}
	if table[curr][w] != 0 {
		return table[curr][w]
	}
	if wt[curr] <= w {
		maxVal1 = val[curr] + recursion(table, n, w-wt[curr], curr+1, val, wt)
	}
	maxVal2 = recursion(table, n, w, curr+1, val, wt)
	table[curr][w] = max(maxVal1, maxVal2)
	return table[curr][w]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
