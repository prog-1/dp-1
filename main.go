package main

import (
	"strconv"
)

func knapsack(w int, val []int, wt []int, cache map[string]int) (max int) {
	if v, ok := cache[strconv.Itoa(w)+" "+strconv.Itoa(len(val))]; ok {
		return v
	}
	for i := 0; i < len(val); i++ {
		if wt[i] <= w {
			tmp := val[i] + knapsack(w-wt[i], val[i+1:], wt[i+1:], cache)
			if tmp > max {
				max = tmp
			}
		}
	}
	cache[strconv.Itoa(w)+" "+strconv.Itoa(len(val))] = max
	return
}
