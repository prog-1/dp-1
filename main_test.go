package main

import (
	"testing"
)

func TestSnackPack(t *testing.T) {
	for _, tc := range []struct {
		n    int
		w    int
		val  []int
		wt   []int
		want int
	}{
		{0, 0, nil, nil, 0},
		{0, 0, []int{}, []int{}, 0},
		{2, 2, []int{0, 0}, []int{0, 0}, 0},
		{1, 9, []int{9}, []int{9}, 9},
		{3, 4, []int{1, 2, 3}, []int{4, 5, 1}, 3},
		{3, 3, []int{1, 2, 3}, []int{4, 5, 6}, 0},
		{4, 7, []int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 22},
		{4, 6, []int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 17},
	} {
		t.Run("", func(t *testing.T) {
			if knapsack(tc.n, tc.w, tc.val, tc.wt) != tc.want {
				t.Errorf("want = %v, got = wrong answer", tc.want)
			}
		})
	}
}
