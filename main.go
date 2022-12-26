package main

func knapsack(leftCap, n int, v, w []int) int {
	if n == 0 || leftCap == 0 {
		return 0
	}
	if w[n-1] > leftCap {
		return knapsack(leftCap, n-1, v, w)
	}
	return max(v[n-1]+knapsack(leftCap-w[n-1], n-1, v, w), knapsack(leftCap, n-1, v, w))
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {
}
