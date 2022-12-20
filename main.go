package main

import "fmt"

func main() {
	// n - item count
	// w - knopsack capacity
	// val[] int - values associated with items
	// wt[] int - item weights
	// find max sum of weights which >= w

	// values := []int{1, 2, 3}
	// weights := []int{4, 5, 6}
	// cap := 3
	// fmt.Println(solveTask(values, weights, cap)) // 0

	// values := []int{2, 3, 1}
	// weights := []int{4, 3, 2}
	// cap := 5
	//fmt.Println(solveTask(values, weights, cap)) // 4

	values := []int{1, 2, 3}
	weights := []int{4, 5, 1}
	cap := 3
	fmt.Println(solveTask(values, weights, cap)) // 3
}

func solveTask(values []int, weights []int, cap int) int {
	// 1. Create all possible combinations
	// 2. Take only those, that <= cap
	// 3. Return the combination with the greatest value
	combs := genCombs(weights, cap)
	var maxValueSum int
	for _, comb := range combs {
		var valueSum int
		for _, el := range comb {
			valueSum += values[find(weights, el)]
		}
		if valueSum > maxValueSum {
			maxValueSum = valueSum
		}
	}
	return maxValueSum
}

func genCombs(s []int, v int) [][]int {
	// How to generate combinations?
	// Take first element
	// Get all without first
	// Create another array with combs from all without first, but with first, and concatenate it with all without first
	if len(s) == 0 {
		return [][]int{{}}
	}
	restCombs := genCombs(s[1:], v)
	var combsWithFirst [][]int
	for _, comb := range restCombs {
		if sum(comb)+s[0] <= v {
			combsWithFirst = append(combsWithFirst, append(comb, s[0]))
		}
	}
	return append(restCombs, combsWithFirst...)
}

func sum(s []int) int {
	var sum int
	for _, el := range s {
		sum += el
	}
	return sum
}

func find(s []int, v int) int {
	for i, el := range s {
		if el == v {
			return i
		}
	}
	panic("must not happen")
}
