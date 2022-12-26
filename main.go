package main

type Key struct {
	n, cap int
}

func knapsack(leftCap, n int, v, w []int, cache map[Key]int) int {
	if res, ok := cache[Key{n, leftCap}]; ok {
		return res
	}
	var res int
	if n == 0 || leftCap == 0 {
		return 0
	}
	if w[n-1] > leftCap {
		res = knapsack(leftCap, n-1, v, w, cache)
	} else {
		res = max(v[n-1]+knapsack(leftCap-w[n-1], n-1, v, w, cache), knapsack(leftCap, n-1, v, w, cache))
	}
	cache[Key{n, leftCap}] = res
	return res
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
