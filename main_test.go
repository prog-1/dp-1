package main

import "testing"

func TestKnapsack(t *testing.T) {
	for _, tc := range []struct {
		name    string
		w       int
		val, wt []int
		want    int
	}{
		{"case-1", 4, []int{1, 2, 3}, []int{4, 5, 1}, 3},
		{"case-2", 3, []int{1, 2, 3}, []int{4, 5, 6}, 0},
		{"case-3", 4, []int{3, 2, 1}, []int{1, 5, 4}, 3},
		{"case-4", 5, []int{1, 2, 3}, []int{4, 5, 1}, 4},
		{"case-5", 6, []int{1, 2, 3}, []int{4, 5, 1}, 5},
		{"case-6", 9, []int{1, 2, 3}, []int{4, 5, 1}, 5},
		{"case-7", 10, []int{1, 2, 3}, []int{4, 5, 1}, 6},
		{"case-8", 11, []int{2, 3, 1, 4, 5}, []int{4, 5, 4, 1, 1}, 14},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := knapsack(tc.w, tc.val, tc.wt, map[string]int{})
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
