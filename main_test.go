package main

import "testing"

func TestIsHeap(t *testing.T) {
	for _, tc := range []struct {
		values, weight []int
		w, n           int
		want           int
	}{
		{nil, nil, 0, 0, 0},
		{[]int{1, 2, 3}, []int{4, 5, 1}, 4, 3, 3},
		{[]int{1, 2, 3}, []int{4, 5, 6}, 3, 3, 0},
		{[]int{1, 2, 3}, []int{4, 5, 6}, 4, 3, 1},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 1, 4, 5, 3}, 100, 6, 21},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 1, 4, 5, 3}, 1, 6, 3},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 1, 4, 5, 3}, 0, 10000000000, 0},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 1, 4, 5, 3}, 100000, 0, 0},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1023, 12309, 210398, 21084, 123, 123}, 13500, 6, 13},
	} {
		if got := knapsack(tc.values, tc.weight, tc.w, tc.n, map[int]int{}); got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
