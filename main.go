package main

import "fmt"

func knapsack(Cap int, V, W []int) int {
	dp := make([][]int, Cap+1)
	for cap := range dp {
		dp[cap] = make([]int, len(V)+1)
	}
	for cap := 1; cap < len(dp); cap++ {
		for item := 1; item < len(dp[cap]); item++ {
			if W[item-1] <= cap {
				dp[cap][item] = max(dp[cap][item-1], dp[cap-W[item-1]][item-1] + V[item-1])
			} else {
				dp[cap][item] = dp[cap][item-1]
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(knapsack(7, []int{2, 2, 4, 5, 3}, []int{3, 1, 3, 4, 2})) // 10
}
