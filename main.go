package main

import "fmt"

func main() {
	// n - item count
	// w - knopsack capacity
	// val[] int - values associated with items
	// wt[] int - item weights
	// find max sum of weights which >= w

	//values := []int{1, 2, 3}
	//weights := []int{4, 5, 6}
	//cap := 3

	//values := []int{2, 3, 1}
	weights := []int{4, 3, 2}
	//cap := 5
	//fmt.Println(solveTask(values, weights, cap)) // 4
	fmt.Println(genCombs(weights))
}

func solveTask(values []int, weights []int, cap int) int {
	// 1. Create all possible combinations
	// 2. Take only those, that <= cap
	// 3. Return the combination with the greatest value
	return 0
}

func genCombs(s []int) [][]int {
	// How to generate combinations?
	// Take first element
	// Get all without first
	// Create another array with combs from all without first, but with first, and concatenate it with all without first
	if len(s) == 0 {
		return [][]int{{}}
	}
	first := s[0]
	rest := s[1:]
	restCombs := genCombs(rest)
	combsWithFirst := make([][]int, len(restCombs))
	for i, comb := range restCombs {
		combsWithFirst[i] = append(comb, first)
	}
	return append(restCombs, combsWithFirst...)
}
