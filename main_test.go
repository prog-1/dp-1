package main

import "testing"

func TestKnapsack2(t *testing.T) {
	for _, tc := range []struct {
		n, w        int
		val, weight []int
		want        int
	}{
		{3, 4, []int{1, 2, 3}, []int{4, 5, 1}, 3},
		{3, 3, []int{1, 2, 3}, []int{4, 5, 6}, 0},
		{9, 120, []int{12, 11, 11, 11, 40, 12, 15, 66, 90}, []int{3, 1, 1, 1, 100, 2, 2, 10, 8}, 228},
		{9, 128, []int{12, 11, 11, 11, 40, 12, 15, 66, 90}, []int{3, 1, 1, 1, 100, 2, 2, 10, 8}, 268},
		{9, 0, []int{12, 11, 11, 11, 40, 12, 15, 66, 90}, []int{3, 1, 1, 1, 100, 2, 2, 10, 8}, 0},
		{10, 8, []int{12, 11, 11, 11, 11, 40, 12, 15, 66, 90}, []int{3, 1, 1, 1, 1, 100, 2, 2, 10, 8}, 90},
	} {
		if got := knapsack2(tc.n, tc.w, tc.val, tc.weight); got != tc.want {
			t.Errorf("Knapsack2(%v, %v, %v, %v) = %v, want = %v", tc.n, tc.w, tc.val, tc.weight, got, tc.want)
		}
	}
}
